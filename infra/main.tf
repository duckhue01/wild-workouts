terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.48.0"
    }

    google = {
      source  = "hashicorp/google"
      version = "4.51.0"
    }

  }
}

provider "aws" {
  region = "ap-southeast-1"
}

provider "google" {
  region = "asia-southeast1"
}

resource "aws_cognito_user_pool" "user" {
  name = "dk-user-pool"
}

resource "aws_cognito_user_pool_client" "client" {
  name         = "client"
  user_pool_id = aws_cognito_user_pool.user.id
}

resource "aws_cognito_user_pool_domain" "customer" {
  domain       = "dk-user-pool-domain"
  user_pool_id = aws_cognito_user_pool.user.id
}

resource "google_project" "gp" {
  name       = "Wild Workouts"
  project_id = "wild-workouts"
}

data "google_client_config" "current" {
}

resource "google_client_oauth_config" "this" {
  provider = google
  name     = va

  client_id                    = google_client_config.this.client_id
  authorized_domains           = ["example.com"]
  redirect_uris                = google_client_config.this.redirect_uris
  javascript_origins           = ["https://example.com"]
  client_secret_length         = 32
  access_type                  = "offline"
  response_types               = ["code"]
  grant_types                  = ["authorization_code", "refresh_token"]
  id_token_signed_response_alg = "RS256"
}
