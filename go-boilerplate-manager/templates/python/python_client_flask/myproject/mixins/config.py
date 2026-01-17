"""Configuration mixin for myproject.

This module provides configuration loading and management capabilities.
"""

import os
from typing import Any, Dict, Optional

from ..types import ConfigDict

class ConfigMixin:
    """Mixin class for configuration management.

    This mixin provides configuration loading from environment variables
    and default configuration setup. It uses cooperative multiple inheritance
    pattern with super() to work properly in the MRO chain.

    Attributes:
        config: ConfigDict instance containing all configuration
    """

    def _load_config(self, custom_config: Optional[Dict[str, Any]] = None) -> ConfigDict:
        """Load configuration from environment variables and custom settings.

        Args:
            custom_config: Custom configuration dictionary to override defaults

        Returns:
            ConfigDict instance with loaded configuration
        """
        custom_config = custom_config or {}

        # Try to load from .env file if python-dotenv is available
        try:
            from dotenv import load_dotenv
            load_dotenv()
        except ImportError:
            pass

        # Load from environment variables with fallbacks
        config = ConfigDict(
            env=custom_config.get('env') or os.getenv('ENV', 'dev'),
            debug=custom_config.get('debug') or os.getenv('DEBUG', 'false').lower() == 'true',
            log_level=custom_config.get('log_level') or os.getenv('LOG_LEVEL', 'INFO'),
            api_key=custom_config.get('api_key') or os.getenv('API_KEY'),
            database_url=custom_config.get('database_url') or os.getenv('DATABASE_URL'),
            timeout=int(custom_config.get('timeout') or os.getenv('TIMEOUT', '30')),
            max_retries=int(custom_config.get('max_retries') or os.getenv('MAX_RETRIES', '3')),
            custom_settings=custom_config.get('custom_settings', {}),
        )

        return config

    def get_config(self, key: str, default: Any = None) -> Any:
        """Get a configuration value.

        Args:
            key: Configuration key to retrieve
            default: Default value if key not found

        Returns:
            Configuration value or default
        """
        return getattr(self.config, key, default)

    def update_config(self, **kwargs: Any) -> None:
        """Update configuration values.

        Args:
            **kwargs: Configuration key-value pairs to update
        """
        for key, value in kwargs.items():
            if hasattr(self.config, key):
                setattr(self.config, key, value)
            else:
                self.config.custom_settings[key] = value

    def print_config(self) -> None:
        """Print current configuration."""
        print("Current Configuration:")
        print("-" * 50)
        for key, value in self.config.to_dict().items():
            # Mask sensitive information
            if key in ['api_key', 'database_url'] and value:
                value = f"{value[:8]}..." if len(str(value)) > 8 else "***"
            print(f"  {key}: {value}")
        print("-" * 50)
