"""Logger mixin for myproject.

This module provides logging functionality.
"""

import os
from datetime import datetime
from typing import Any, Optional

from ..types import Result


class LoggerMixin:
    """Mixin class for logging operations.

    Provides logging functionality with different log levels.
    Uses cooperative multiple inheritance pattern with super().

    Attributes:
        log_level: Current logging level
        log_file: Path to log file (optional)
        log_enabled: Whether logging is enabled
    """

    LOG_LEVELS = {
        "DEBUG": 10,
        "INFO": 20,
        "WARNING": 30,
        "ERROR": 40,
        "CRITICAL": 50,
    }

    def _should_log(self, level: str) -> bool:
        """Check if a message at the given level should be logged.

        Args:
            level: Log level to check

        Returns:
            True if the message should be logged
        """
        if not self.log_enabled:
            return False

        current_level = self.LOG_LEVELS.get(self.log_level.upper(), 20)
        message_level = self.LOG_LEVELS.get(level.upper(), 20)

        return message_level >= current_level

    def _format_log_message(self, level: str, message: str, context: Optional[dict] = None) -> str:
        """Format a log message.

        Args:
            level: Log level
            message: Log message
            context: Additional context data

        Returns:
            Formatted log message
        """
        timestamp = datetime.utcnow().isoformat()
        formatted = f"[{timestamp}] [{level.upper()}] {message}"

        if context:
            formatted += f" | Context: {context}"

        return formatted

    def _write_log(self, formatted_message: str) -> None:
        """Write a log message to output.

        Args:
            formatted_message: The formatted log message
        """
        # Write to console
        print(formatted_message)

        # Write to file if configured
        if self.log_file:
            try:
                with open(self.log_file, 'a', encoding='utf-8') as f:
                    f.write(formatted_message + '\n')
            except Exception as e:
                print(f"Failed to write to log file: {e}")

    def log(self, level: str, message: str, context: Optional[dict] = None) -> Result:
        """Log a message at the specified level.

        Args:
            level: Log level (DEBUG, INFO, WARNING, ERROR, CRITICAL)
            message: Message to log
            context: Additional context information

        Returns:
            Result dictionary with log status
        """
        try:
            if not self._should_log(level):
                return {
                    "success": True,
                    "data": {"message": "Log level filtered", "logged": False},
                }

            formatted_message = self._format_log_message(level, message, context)
            self._write_log(formatted_message)

            return {
                "success": True,
                "data": {
                    "message": "Logged successfully",
                    "logged": True,
                    "level": level,
                },
            }
        except Exception as e:
            return {
                "success": False,
                "error": f"Logging failed: {str(e)}",
                "data": None,
            }

    def debug(self, message: str, context: Optional[dict] = None) -> Result:
        """Log a debug message.

        Args:
            message: Message to log
            context: Additional context

        Returns:
            Result dictionary
        """
        return self.log("DEBUG", message, context)

    def info(self, message: str, context: Optional[dict] = None) -> Result:
        """Log an info message.

        Args:
            message: Message to log
            context: Additional context

        Returns:
            Result dictionary
        """
        return self.log("INFO", message, context)

    def warning(self, message: str, context: Optional[dict] = None) -> Result:
        """Log a warning message.

        Args:
            message: Message to log
            context: Additional context

        Returns:
            Result dictionary
        """
        return self.log("WARNING", message, context)

    def error(self, message: str, context: Optional[dict] = None) -> Result:
        """Log an error message.

        Args:
            message: Message to log
            context: Additional context

        Returns:
            Result dictionary
        """
        return self.log("ERROR", message, context)

    def critical(self, message: str, context: Optional[dict] = None) -> Result:
        """Log a critical message.

        Args:
            message: Message to log
            context: Additional context

        Returns:
            Result dictionary
        """
        return self.log("CRITICAL", message, context)

    def set_log_level(self, level: str) -> Result:
        """Set the logging level.

        Args:
            level: New log level

        Returns:
            Result dictionary
        """
        if level.upper() not in self.LOG_LEVELS:
            return {
                "success": False,
                "error": f"Invalid log level: {level}. Valid levels: {', '.join(self.LOG_LEVELS.keys())}",
                "data": None,
            }

        old_level = self.log_level
        self.log_level = level.upper()

        return {
            "success": True,
            "data": {
                "message": "Log level updated",
                "old_level": old_level,
                "new_level": self.log_level,
            },
        }
