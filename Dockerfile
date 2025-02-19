FROM node:14

# Set the working directory
WORKDIR /usr/src/app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application files
COPY . .

# Rebuild native modules
RUN npm rebuild sqlite3

# Expose the port the app runs on
EXPOSE 3000

# Create a volume for the SQLite database
VOLUME ["/usr/src/app/src/database"]

# Command to run the application
CMD ["node", "src/index.js"]