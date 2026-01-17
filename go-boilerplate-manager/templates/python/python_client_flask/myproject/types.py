"""Type definitions for myproject.

This module contains all TypedDict and dataclass definitions used throughout the project.
"""

from dataclasses import dataclass, field
from typing import Any, Dict, Optional, TypedDict


class Result(TypedDict, total=False):
    """Result type for operation returns.

    Attributes:
        success: Whether the operation was successful
        data: The result data
        error: Error message if operation failed
        metadata: Additional metadata about the operation
    """
    success: bool
    data: Any
    error: Optional[str]
    metadata: Dict[str, Any]


@dataclass
class ConfigDict:
    """Configuration dictionary for the application.

    Attributes:
        env: Environment name (dev, staging, prod)
        debug: Enable debug mode
        log_level: Logging level
        api_key: API key for external services
        database_url: Database connection URL
        timeout: Request timeout in seconds
        max_retries: Maximum number of retries for failed operations
        custom_settings: Additional custom settings
    """
    env: str = "dev"
    debug: bool = False
    log_level: str = "INFO"
    api_key: Optional[str] = None
    database_url: Optional[str] = None
    timeout: int = 30
    max_retries: int = 3
    custom_settings: Dict[str, Any] = field(default_factory=dict)

    def to_dict(self) -> Dict[str, Any]:
        """Convert config to dictionary."""
        return {
            "env": self.env,
            "debug": self.debug,
            "log_level": self.log_level,
            "api_key": self.api_key,
            "database_url": self.database_url,
            "timeout": self.timeout,
            "max_retries": self.max_retries,
            "custom_settings": self.custom_settings,
        }
