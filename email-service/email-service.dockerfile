
# Use the official Node.js image as the base image
FROM node:12.16.1

# Set the working directory in the container
WORKDIR /app

# Copy the package.json and package-lock.json files to the container
COPY package*.json ./

# Install the dependencies
RUN npm install

# Copy the application code to the container
COPY . .

EXPOSE 50051

# Start the application
CMD ["npm", "start"]
