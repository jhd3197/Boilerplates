from setuptools import setup, find_packages
import os

# Read metadata without importing the package
metadata = {}
with open(os.path.join("myproject", "metadata.py"), encoding="utf-8") as f:
    exec(f.read(), metadata)

setup(
    name=metadata["NAME"],
    version=metadata["VERSION"],
    description=metadata["DESCRIPTION"],
    author=metadata["AUTHOR"],
    packages=find_packages(),
    install_requires=[
        "flask",
        "python-dotenv",
        "requests", # Added requests as it is often needed for clients
    ],
    python_requires=">=3.8",
)
