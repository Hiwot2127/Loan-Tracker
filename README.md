# Loan Tracker API

## Project Overview

This is a Loan Tracker API built using Golang and the Gin framework. The API allows users to manage loans and user accounts, with both user and admin functionalities. The project follows clean architecture principles to ensure maintainability and scalability.

## Features

- **User Registration:** Allows users to register with an email, password, and profile details.
- **Email Verification:** Users can verify their email via a token sent to their email address.
- **User Login:** Authenticate users and provide access and refresh tokens.
- **User Profile:** Retrieve authenticated user profile details.
- **Password Reset:** Request a password reset and update the password using a token.
- **Admin Functions:** Admins can view all users and delete user accounts.

## Endpoints

### User Management

- **POST /users/register:** Register a new user.
- **GET /users/verify-email:** Verify a user's email address.
- **POST /users/login:** Login and obtain tokens.
- **GET /users/profile:** Get the authenticated user's profile.
- **POST /users/password-reset:** Request a password reset.
- **POST /users/password-reset/confirm:** Confirm password reset with a token.

### Admin Functions

- **GET /admin/users:** Retrieve a list of all users.
- **DELETE /admin/users/{id}:** Delete a specific user account.