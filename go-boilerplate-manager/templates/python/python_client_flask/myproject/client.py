"""Client class for MyProject.

This module contains the main Client class that aggregates all mixins
using cooperative multiple inheritance.
"""

import pathlib
import gc
import os
from typing import Any, Dict, Optional

from dotenv import load_dotenv

from .mixins.config import ConfigMixin
from .mixins.tools import ToolsMixin
from .mixins.database import DatabaseMixin
from .mixins.api_client import APIMixin
from .mixins.logger import LoggerMixin
from .mixins.cache import CacheMixin
from .types import Result

# Load environment variables from .env file
load_dotenv()


class Client(
    CacheMixin,
    LoggerMixin,
    APIMixin,
    DatabaseMixin,
    ToolsMixin,
    ConfigMixin,
):
    """Client class that combines all mixins using cooperative inheritance.

    This class demonstrates the proper use of multiple inheritance in Python,
    combining all available mixins. The MRO (Method Resolution Order) ensures
    that all __init__ methods are called in the correct order using super().

    The MRO for this class is:
        Client -> CacheMixin -> LoggerMixin -> APIMixin -> DatabaseMixin ->
        ToolsMixin -> ConfigMixin -> object

    Attributes:
        name: Name of this Client instance
        version: Version of the application
        JS_FRAMEWORK: JavaScript framework identifier (class variable)

    All mixin attributes are also available:
        - config (from ConfigMixin)
        - cache_folder, cache_storage (from CacheMixin)
        - log_level, log_file (from LoggerMixin)
        - api_base_url, api_key (from APIMixin)
        - db_url, db_connection (from DatabaseMixin)

    Example:
        >>> client = Client(name="MyApp", config={"debug": True})
        >>> client.info("Application started")
        >>> result = client.calculate("add", 1, 2, 3)
        >>> client.cache_set("result", result["data"])
    """

    # Class variables
    JS_FRAMEWORK = ""
    DEFAULT_CACHE_FOLDER = "cache"
    DEFAULT_VERSION = "1.0.0"

    def __init__(
        self,
        name: str = "MyProject",
        version: Optional[str] = None,
        **kwargs: Any,
    ) -> None:
        """Initialize the Client instance.

        Args:
            name: Name of the application
            version: Version of the application (uses DEFAULT_VERSION if None)
            **kwargs: Additional keyword arguments
                - config: Configuration dictionary
                - cache_folder: Cache folder path
                - log_level: Logging level
                - api_base_url: API base URL
                - db_url: Database URL
                - etc.
        """
        # Initialize object
        super().__init__()

        # --- ConfigMixin Initialization ---
        custom_config = kwargs.get('config', {})
        # _load_config is available from ConfigMixin
        self.config = self._load_config(custom_config)

        # --- CacheMixin Initialization ---
        # Get cache settings from config or environment
        self.cache_folder = kwargs.get('cache_folder') or os.getenv('CACHE_FOLDER', 'cache')
        self.cache_ttl = kwargs.get('cache_ttl', 3600)  # 1 hour default
        
        # Initialize in-memory cache
        self.cache_storage: dict[str, dict] = {}
        
        # Create cache folder if it doesn't exist
        pathlib.Path(self.cache_folder).mkdir(parents=True, exist_ok=True)

        # --- LoggerMixin Initialization ---
        self.log_level = kwargs.get('log_level') or os.getenv('LOG_LEVEL', 'INFO')
        self.log_file = kwargs.get('log_file') or os.getenv('LOG_FILE')
        self.log_enabled = kwargs.get('log_enabled', True)

        if hasattr(self, 'config'):
            self.log_level = self.config.log_level

        # --- APIMixin Initialization ---
        self.api_base_url = kwargs.get('api_base_url') or os.getenv('API_BASE_URL', '')
        self.api_key = kwargs.get('api_key') or os.getenv('API_KEY')
        self.api_timeout = kwargs.get('api_timeout', 30)

        if hasattr(self, 'config'):
            if self.config.api_key:
                self.api_key = self.config.api_key
            self.api_timeout = self.config.timeout

        # --- DatabaseMixin Initialization ---
        self.db_url = kwargs.get('db_url') or os.getenv('DATABASE_URL')
        self.db_connection = None

        if hasattr(self, 'config') and self.config.database_url:
            self.db_url = self.config.database_url

        # --- Client Initialization ---
        # Set version with fallback
        version = version or os.getenv('APP_VERSION', self.DEFAULT_VERSION)
        
        # Set Client-specific attributes
        self.name = name
        self.version = version

        # Run garbage collection
        gc.collect()

        # Log initialization if debug mode is enabled
        if hasattr(self, 'config') and self.config.debug:
            self.debug(
                f"{self.name} v{self.version} initialized",
                {"cache_folder": self.cache_folder}
            )

    def initialize(self) -> Result:
        """Initialize all services and connections.

        This method sets up database connections, API clients, and other
        services that require initialization after the Client is created.

        Returns:
            Result dictionary with initialization status
        """
        try:
            results = []

            # Connect to database if URL is configured
            if self.db_url:
                db_result = self.connect_db()
                results.append(("database", db_result["success"]))
                if db_result["success"]:
                    self.info("Database connected successfully")
                else:
                    self.warning(f"Database connection failed: {db_result.get('error', 'Unknown error')}")

            # Log cache stats
            cache_stats = self.cache_stats()
            if cache_stats["success"]:
                self.debug("Cache initialized", cache_stats["data"])

            self.info(f"{self.name} v{self.version} - All services initialized")

            return {
                "success": True,
                "data": {
                    "message": "Client initialized successfully",
                    "name": self.name,
                    "version": self.version,
                    "services": dict(results),
                },
            }
        except Exception as e:
            self.error(f"Initialization failed: {str(e)}")
            return {
                "success": False,
                "error": f"Initialization failed: {str(e)}",
                "data": None,
            }

    def shutdown(self) -> Result:
        """Shutdown all services and cleanup resources.

        Returns:
            Result dictionary with shutdown status
        """
        try:
            # Disconnect from database
            if self.db_connection:
                db_result = self.disconnect_db()
                if db_result["success"]:
                    self.info("Database disconnected")

            # Clear cache if configured
            clear_cache_on_shutdown = os.getenv('CLEAR_CACHE_ON_SHUTDOWN', 'false').lower() == 'true'
            if clear_cache_on_shutdown:
                cache_result = self.cache_clear(clear_disk=True)
                if cache_result["success"]:
                    self.info("Cache cleared on shutdown")

            self.info(f"{self.name} shutdown complete")

            # Final garbage collection
            gc.collect()

            return {
                "success": True,
                "data": {"message": "Client shutdown successfully"},
            }
        except Exception as e:
            self.error(f"Shutdown failed: {str(e)}")
            return {
                "success": False,
                "error": f"Shutdown failed: {str(e)}",
                "data": None,
            }

    def status(self) -> Result:
        """Get the current status of the Client and all services.

        Returns:
            Result dictionary with status information
        """
        try:
            cache_stats = self.cache_stats()

            status_info = {
                "name": self.name,
                "version": self.version,
                "config": self.config.to_dict() if hasattr(self, 'config') else {},
                "database_connected": self.db_connection is not None,
                "cache_stats": cache_stats["data"] if cache_stats["success"] else {},
                "log_level": self.log_level if hasattr(self, 'log_level') else "INFO",
            }

            return {
                "success": True,
                "data": status_info,
            }
        except Exception as e:
            return {
                "success": False,
                "error": f"Failed to get status: {str(e)}",
                "data": None,
            }

    def get_info(self) -> Dict[str, Any]:
        """Get detailed information about this Client instance.

        Returns:
            Dictionary containing instance information including MRO
        """
        return {
            "name": self.name,
            "version": self.version,
            "mro": [cls.__name__ for cls in type(self).__mro__],
            "config": self.config.to_dict() if hasattr(self, 'config') else {},
            "available_mixins": [
                "CacheMixin",
                "LoggerMixin",
                "APIMixin",
                "DatabaseMixin",
                "ToolsMixin",
                "ConfigMixin",
            ],
            "js_framework": self.JS_FRAMEWORK,
        }

    def __repr__(self) -> str:
        """Return string representation of Client instance."""
        return f"<Client(name='{self.name}', version='{self.version}')>"

    def __str__(self) -> str:
        """Return human-readable string of Client instance."""
        return f"{self.name} v{self.version}"

    def __enter__(self):
        """Context manager entry."""
        self.initialize()
        return self

    def __exit__(self, exc_type, exc_val, exc_tb):
        """Context manager exit."""
        self.shutdown()
        return False
