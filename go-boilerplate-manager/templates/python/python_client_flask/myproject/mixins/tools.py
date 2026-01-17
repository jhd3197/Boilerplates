"""Tools mixin for myproject.

This module provides utility methods and helper functions.
"""

import json
from datetime import datetime
from typing import Any, Dict, List, Optional

from ..types import Result


class ToolsMixin:
    """Mixin class providing utility tools and helper methods.

    This mixin provides various utility methods for data processing,
    formatting, and common operations. It uses cooperative multiple
    inheritance pattern with super() to work properly in the MRO chain.
    """

    def calculate(self, operation: str, *values: float) -> Result:
        """Perform a calculation on the provided values.

        Args:
            operation: The operation to perform (add, subtract, multiply, divide, average)
            *values: Numeric values to calculate

        Returns:
            Result dictionary with success status and calculated value
        """
        if not values:
            return {
                "success": False,
                "error": "No values provided for calculation",
                "data": None,
            }

        try:
            if operation == "add":
                result = sum(values)
            elif operation == "subtract":
                result = values[0] - sum(values[1:])
            elif operation == "multiply":
                result = 1
                for val in values:
                    result *= val
            elif operation == "divide":
                result = values[0]
                for val in values[1:]:
                    if val == 0:
                        return {
                            "success": False,
                            "error": "Division by zero",
                            "data": None,
                        }
                    result /= val
            elif operation == "average":
                result = sum(values) / len(values)
            else:
                return {
                    "success": False,
                    "error": f"Unknown operation: {operation}",
                    "data": None,
                }

            return {
                "success": True,
                "data": result,
                "metadata": {
                    "operation": operation,
                    "values_count": len(values),
                    "timestamp": datetime.utcnow().isoformat(),
                },
            }
        except Exception as e:
            return {
                "success": False,
                "error": str(e),
                "data": None,
            }

    def format_data(self, data: Any, format_type: str = "json") -> Result:
        """Format data into different representations.

        Args:
            data: Data to format
            format_type: Output format (json, pretty_json, str, repr)

        Returns:
            Result dictionary with formatted data
        """
        try:
            if format_type == "json":
                formatted = json.dumps(data)
            elif format_type == "pretty_json":
                formatted = json.dumps(data, indent=2, sort_keys=True)
            elif format_type == "str":
                formatted = str(data)
            elif format_type == "repr":
                formatted = repr(data)
            else:
                return {
                    "success": False,
                    "error": f"Unknown format type: {format_type}",
                    "data": None,
                }

            return {
                "success": True,
                "data": formatted,
                "metadata": {
                    "format_type": format_type,
                    "original_type": type(data).__name__,
                },
            }
        except Exception as e:
            return {
                "success": False,
                "error": f"Failed to format data: {str(e)}",
                "data": None,
            }

    def validate_data(self, data: Dict[str, Any], required_fields: List[str]) -> Result:
        """Validate that data contains all required fields.

        Args:
            data: Dictionary to validate
            required_fields: List of required field names

        Returns:
            Result dictionary with validation status
        """
        if not isinstance(data, dict):
            return {
                "success": False,
                "error": "Data must be a dictionary",
                "data": None,
            }

        missing_fields = [field for field in required_fields if field not in data]

        if missing_fields:
            return {
                "success": False,
                "error": f"Missing required fields: {', '.join(missing_fields)}",
                "data": None,
                "metadata": {"missing_fields": missing_fields},
            }

        return {
            "success": True,
            "data": data,
            "metadata": {
                "validated_fields": required_fields,
                "total_fields": len(data),
            },
        }

    def transform_dict(
        self,
        data: Dict[str, Any],
        key_transform: Optional[str] = None,
        value_transform: Optional[callable] = None,
    ) -> Result:
        """Transform dictionary keys and/or values.

        Args:
            data: Dictionary to transform
            key_transform: Transformation to apply to keys (upper, lower, title)
            value_transform: Function to apply to values

        Returns:
            Result dictionary with transformed data
        """
        try:
            transformed = {}

            for key, value in data.items():
                # Transform key
                if key_transform == "upper":
                    new_key = key.upper()
                elif key_transform == "lower":
                    new_key = key.lower()
                elif key_transform == "title":
                    new_key = key.title()
                else:
                    new_key = key

                # Transform value
                if value_transform and callable(value_transform):
                    new_value = value_transform(value)
                else:
                    new_value = value

                transformed[new_key] = new_value

            return {
                "success": True,
                "data": transformed,
                "metadata": {
                    "key_transform": key_transform,
                    "value_transform_applied": value_transform is not None,
                },
            }
        except Exception as e:
            return {
                "success": False,
                "error": f"Transformation failed: {str(e)}",
                "data": None,
            }
