"""Playground for testing MyProject.

This module demonstrates the usage of the MyProject,
showcasing the Client class, mixins, and the new Flask API.
"""

from myproject.client import Client
import sys

def separator(title: str) -> None:
    """Print a visual separator with title."""
    print("\n" + "=" * 70)
    print(f"  {title}")
    print("=" * 70 + "\n")


def test_basic_initialization():
    """Test basic Client initialization."""
    separator("1. Basic Client Initialization")

    # Create instance with default settings
    client = Client()
    print(f"Created: {client}")
    print(f"Repr: {repr(client)}")
    print(f"MRO: {' -> '.join(client.get_info()['mro'][:7])}")


def test_api_client_mixin():
    """Test the API Client Mixin (formerly api.py)."""
    separator("2. API Client Mixin (Client -> External APIs)")
    
    print("This demonstrates the Client making requests TO external services.")
    print("Note: The mixin is now in 'mixins/api_client.py' to avoid conflict")
    print("with the 'api/' folder which contains the Flask server.\n")

    client = Client(
        name="APIDemo",
        api_base_url="https://api.example.com",
        api_key="test_api_key_12345"
    )

    # Test GET request via client
    result = client.get("/users", params={"page": 1, "limit": 10})
    print(f"Client GET request: {result['data']}")


def test_flask_server_structure():
    """Explain the new Flask Server structure."""
    separator("3. Flask API Server Structure")
    
    print("The project now includes a Flask API server in the 'api/' directory.")
    print("Refactored structure:")
    print("  app.py                 -> Entry point to run the server")
    print("  myproject/api/         -> Flask App Package")
    print("  myproject/api/blueprints/ -> Route Blueprints")
    print("  myproject/mixins/      -> Client Mixins (Logic, DB, Tools, etc.)")
    print("\nTo run the server:")
    print("  $ python app.py")
    print("\nTry visiting: http://localhost:3167/api/v1/health")


def main():
    """Run playground tests."""
    print("\n" + "#" * 70)
    print("#" + " " * 68 + "#")
    print("#" + "  MyProject Boilerplate - Playground".center(68) + "#")
    print("#" + " " * 68 + "#")
    print("#" * 70)

    test_basic_initialization()
    test_api_client_mixin()
    test_flask_server_structure()

    print("\n" + "=" * 70)
    print("Summary:")
    print("  - Client structure remains intact (mixins/ folder)")
    print("  - API Client mixin renamed to api_client.py")
    print("  - New Flask API added in api/ folder")
    print("  - New app.py entry point created")
    print("=" * 70 + "\n")

if __name__ == "__main__":
    main()
