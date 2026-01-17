from myproject.api import create_app
from myproject.client import Client
from dotenv import load_dotenv

# Load environment variables
load_dotenv()

# Create Flask app instance
app = create_app()

import os

if __name__ == "__main__":
    # You can instantiate the Client here if needed for background tasks or config
    client = Client()
    client.info("Starting Flask API Server...")
    
    # Run the server
    port = int(os.environ.get("PORT", 3167))
    app.run(debug=True, host='0.0.0.0', port=port)
