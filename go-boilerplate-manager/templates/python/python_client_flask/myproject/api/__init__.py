from flask import Flask
from .blueprints import health_bp
from .. import metadata
from datetime import datetime, timedelta

def create_app(config_object=None):
    app = Flask(__name__)
    start_time = datetime.now()
    
    if config_object:
        app.config.from_object(config_object)
    
    # Register Blueprints
    app.register_blueprint(health_bp, url_prefix='/api/v1')

    @app.route('/')
    def index():
        uptime = datetime.now() - start_time
        # Remove microseconds for cleaner output
        uptime_str = str(uptime).split('.')[0]
        
        return {
            "name": metadata.NAME,
            "description": metadata.DESCRIPTION,
            "version": metadata.VERSION,
            "author": metadata.AUTHOR,
            "uptime": uptime_str,
            "endpoints": {
                "health_check": "/api/v1/health"
            }
        }
    
    return app
