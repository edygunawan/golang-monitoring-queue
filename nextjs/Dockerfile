# Use Node.js 22 LTS as the base image
FROM node:22-bullseye

# Set working directory inside the container
WORKDIR /app/dashboard


# Copy package.json and package-lock.json
# COPY ./package.json ./package-lock.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
# COPY . .

# Expose the port used by the Next.js application
EXPOSE 3000

# Command to start the development server
CMD ["npm", "run", "dev"]
