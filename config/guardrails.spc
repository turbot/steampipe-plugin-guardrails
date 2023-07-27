connection "guardrails" {
  plugin = "guardrails"

  # Default credentials:
  # By default, Steampipe will use your Turbot Guardrails profiles and credentials exactly
  # the same as the Turbot Guardrails CLI and Turbot Guardrails Terraform provider. In many cases, no
  # extra configuration is required to use Steampipe.

  # Use an existing Turbot profile configured in ~/.config/turbot
  # profile = "my-profile"

  # Define exact connection parameters to Turbot. This takes precedence over all
  # Turbot configuration, profile and environment variables.
  # This can also be also be set using TURBOT_ACCESS_KEY, TURBOT_SECRET_KEY and TURBOT_WORKSPACE env variables.
  # workspace  = "https://turbot-acme.cloud.turbot.com/"
  # access_key = "c8e2c2ed-1ca8-429b-b369-010e3cf75aac"
  # secret_key = "a3d8385d-47f7-40c5-a90c-bfdf5b43c8dd"
}
