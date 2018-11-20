# GOCI-Admin

This app will scan your OCI env and will stop your instances after a certain time

# Prereqs
 - Must have a $HOME/.oci/config configured for your tenancy
 - Have an the $C environment variable declared with the compartment that you want to scan for
 ```
 export c=<compartment-id>
 ```