#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <netinet/in.h>
#include <signal.h>

int server_fd;

void handle_sigint(int sig) {
    printf("\nShutting down server...\n");
    if (server_fd >= 0) {
        close(server_fd);
        printf("Socket closed.\n");
    }
    exit(0);
}

int main() {
    struct sockaddr_in address;
    int addrlen = sizeof(address);
    int client_fd;
    char buffer[4096] = {0};

    const char* http_response =
        "HTTP/1.1 200 OK\r\n"
        "Content-Type: text/plain\r\n"
        "Content-Length: 12\r\n"
        "\r\n"
        "Hello World!";

    // Register signal handler
    signal(SIGINT, handle_sigint);

    // Create socket
    if ((server_fd = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
        perror("Socket failed");
        exit(EXIT_FAILURE);
    }

    // Allow address reuse
    int opt = 1;
    setsockopt(server_fd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt));

    // Bind socket
    address.sin_family = AF_INET;
    address.sin_addr.s_addr = INADDR_ANY;
    address.sin_port = htons(8083);

    if (bind(server_fd, (struct sockaddr*)&address, sizeof(address)) < 0) {
        perror("Bind failed");
        close(server_fd);
        exit(EXIT_FAILURE);
    }

    // Listen
    if (listen(server_fd, 10) < 0) {
        perror("Listen failed");
        close(server_fd);
        exit(EXIT_FAILURE);
    }

    printf("Server listening on port 8083. Press Ctrl+C to stop.\n");

    // Accept loop
    while (1) {
        if ((client_fd = accept(server_fd, (struct sockaddr*)&address, (socklen_t*)&addrlen)) < 0) {
            perror("Accept failed");
            continue;
        }

        // Read request
        read(client_fd, buffer, sizeof(buffer) - 1);
        printf("Request:\n%s\n", buffer);

        // Respond
        send(client_fd, http_response, strlen(http_response), 0);

        // Close client
        close(client_fd);
    }

    // Should never reach here due to infinite loop, but just in case:
    close(server_fd);
    return 0;
}
