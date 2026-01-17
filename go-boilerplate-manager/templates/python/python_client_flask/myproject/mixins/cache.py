"""Cache mixin for myproject.

This module provides caching functionality.
"""

import os
import pathlib
import json
import time
from typing import Any, Optional

from ..types import Result

class CacheMixin:
    """Mixin class for caching operations.

    Provides in-memory and file-based caching functionality.
    Uses cooperative multiple inheritance pattern with super().

    Attributes:
        cache_folder: Path to cache folder
        cache_storage: In-memory cache storage
        cache_ttl: Default time-to-live for cache entries (seconds)
    """

    def _get_cache_file_path(self, key: str) -> str:
        """Get the file path for a cache key.

        Args:
            key: Cache key

        Returns:
            Path to cache file
        """
        # Simple sanitization of key for filename
        safe_key = "".join(c if c.isalnum() or c in "._-" else "_" for c in key)
        return os.path.join(self.cache_folder, f"{safe_key}.cache")

    def _is_expired(self, timestamp: float, ttl: Optional[int] = None) -> bool:
        """Check if a cache entry is expired.

        Args:
            timestamp: Cache entry timestamp
            ttl: Time-to-live in seconds (uses default if None)

        Returns:
            True if expired
        """
        ttl = ttl if ttl is not None else self.cache_ttl
        return (time.time() - timestamp) > ttl

    def cache_set(
        self,
        key: str,
        value: Any,
        ttl: Optional[int] = None,
        persist: bool = False,
    ) -> Result:
        """Set a value in cache.

        Args:
            key: Cache key
            value: Value to cache
            ttl: Time-to-live in seconds (uses default if None)
            persist: Whether to persist to disk

        Returns:
            Result dictionary with cache status
        """
        try:
            cache_entry = {
                "value": value,
                "timestamp": time.time(),
                "ttl": ttl if ttl is not None else self.cache_ttl,
            }

            # Store in memory
            self.cache_storage[key] = cache_entry

            # Optionally persist to disk
            if persist:
                cache_file = self._get_cache_file_path(key)
                with open(cache_file, 'w', encoding='utf-8') as f:
                    json.dump(cache_entry, f)

            return {
                "success": True,
                "data": {
                    "message": "Value cached successfully",
                    "key": key,
                    "persisted": persist,
                },
            }
        except Exception as e:
            return {
                "success": False,
                "error": f"Failed to cache value: {str(e)}",
                "data": None,
            }

    def cache_get(self, key: str, from_disk: bool = False) -> Result:
        """Get a value from cache.

        Args:
            key: Cache key
            from_disk: Whether to load from disk if not in memory

        Returns:
            Result dictionary with cached value
        """
        try:
            # Try to get from memory first
            if key in self.cache_storage:
                cache_entry = self.cache_storage[key]

                # Check if expired
                if self._is_expired(cache_entry["timestamp"], cache_entry["ttl"]):
                    del self.cache_storage[key]
                    return {
                        "success": False,
                        "error": "Cache entry expired",
                        "data": None,
                    }

                return {
                    "success": True,
                    "data": cache_entry["value"],
                    "metadata": {
                        "source": "memory",
                        "timestamp": cache_entry["timestamp"],
                    },
                }

            # Try to load from disk if requested
            if from_disk:
                cache_file = self._get_cache_file_path(key)
                if os.path.exists(cache_file):
                    with open(cache_file, 'r', encoding='utf-8') as f:
                        cache_entry = json.load(f)

                    # Check if expired
                    if self._is_expired(cache_entry["timestamp"], cache_entry["ttl"]):
                        os.remove(cache_file)
                        return {
                            "success": False,
                            "error": "Cache entry expired",
                            "data": None,
                        }

                    # Load into memory
                    self.cache_storage[key] = cache_entry

                    return {
                        "success": True,
                        "data": cache_entry["value"],
                        "metadata": {
                            "source": "disk",
                            "timestamp": cache_entry["timestamp"],
                        },
                    }

            return {
                "success": False,
                "error": "Cache key not found",
                "data": None,
            }
        except Exception as e:
            return {
                "success": False,
                "error": f"Failed to get cached value: {str(e)}",
                "data": None,
            }

    def cache_delete(self, key: str, from_disk: bool = True) -> Result:
        """Delete a value from cache.

        Args:
            key: Cache key
            from_disk: Whether to also delete from disk

        Returns:
            Result dictionary with deletion status
        """
        try:
            deleted_from = []

            # Delete from memory
            if key in self.cache_storage:
                del self.cache_storage[key]
                deleted_from.append("memory")

            # Delete from disk
            if from_disk:
                cache_file = self._get_cache_file_path(key)
                if os.path.exists(cache_file):
                    os.remove(cache_file)
                    deleted_from.append("disk")

            if deleted_from:
                return {
                    "success": True,
                    "data": {
                        "message": "Cache entry deleted",
                        "deleted_from": deleted_from,
                    },
                }
            else:
                return {
                    "success": False,
                    "error": "Cache key not found",
                    "data": None,
                }
        except Exception as e:
            return {
                "success": False,
                "error": f"Failed to delete cache entry: {str(e)}",
                "data": None,
            }

    def cache_clear(self, clear_disk: bool = False) -> Result:
        """Clear all cache entries.

        Args:
            clear_disk: Whether to also clear disk cache

        Returns:
            Result dictionary with clear status
        """
        try:
            # Clear memory cache
            memory_count = len(self.cache_storage)
            self.cache_storage.clear()

            disk_count = 0
            # Clear disk cache if requested
            if clear_disk:
                cache_files = pathlib.Path(self.cache_folder).glob("*.cache")
                for cache_file in cache_files:
                    cache_file.unlink()
                    disk_count += 1

            return {
                "success": True,
                "data": {
                    "message": "Cache cleared",
                    "memory_entries_cleared": memory_count,
                    "disk_entries_cleared": disk_count,
                },
            }
        except Exception as e:
            return {
                "success": False,
                "error": f"Failed to clear cache: {str(e)}",
                "data": None,
            }

    def cache_stats(self) -> Result:
        """Get cache statistics.

        Returns:
            Result dictionary with cache stats
        """
        try:
            # Count disk cache files
            cache_files = list(pathlib.Path(self.cache_folder).glob("*.cache"))
            disk_count = len(cache_files)

            # Calculate total disk size
            disk_size = sum(f.stat().st_size for f in cache_files)

            return {
                "success": True,
                "data": {
                    "memory_entries": len(self.cache_storage),
                    "disk_entries": disk_count,
                    "disk_size_bytes": disk_size,
                    "cache_folder": self.cache_folder,
                    "default_ttl": self.cache_ttl,
                },
            }
        except Exception as e:
            return {
                "success": False,
                "error": f"Failed to get cache stats: {str(e)}",
                "data": None,
            }
