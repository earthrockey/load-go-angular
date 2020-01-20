# Use the standard Nginx image from Docker Hub
FROM nginx

# The Dockerfile's author
LABEL Usama Ashraf

# Copy custom configuration file from the current directory
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

# Start Nginx when the container has provisioned.
CMD ["nginx", "-g", "daemon off;"]
