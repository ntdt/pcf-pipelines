#!/bin/bash

set -eu
 
properties=$(
    jq -n \
       --arg nsx_api_managers "${NSX_API_MANAGERS}" \
       --arg nsx_auth 'simple' \
       --arg nsx_api_user "${NSX_API_USER}" \
       --arg nsx_api_password "${NSX_API_PASSWORD}" \
       --arg nsx_api_ca_cert "${NSX_API_CA_CERT}" \
       --arg foundation_name "${NCP_FOUNDATION_NAME}" \
       --arg overlay_tz "${NCP_OVERLAY_TZ}" \
       --arg tier0_router "${NCP_TIER0_ROUTER}" \
       --arg container_block_name "${NCP_CONTAINER_BLOCK_NAME}" \
       --arg subnet_prefix "${NCP_SUBNET_PREFIX}" \
    '
    {
       ".properties.nsx_api_managers": {
          "value": $nsx_api_managers
       },
       ".properties.nsx_auth": {
          "value": $nsx_auth
       },
       ".properties.nsx_auth.simple.nsx_api_user": {
          "value": $nsx_api_user
       },
       ".properties.nsx_auth.simple.nsx_api_password": {
          "value": {
            "secret": $nsx_api_password
          }
       },
       ".properties.nsx_api_ca_cert": {
          "value": $nsx_api_ca_cert
       },
       ".properties.foundation_name": {
          "value": $foundation_name
       },
       ".properties.overlay_tz": {
          "value": $overlay_tz
       },
       ".properties.tier0_router": {
          "value": $tier0_router
       },
       ".properties.container_ip_blocks": {
         "value": [
           {
             "name": $container_block_name
           }
         ]
       },
       ".properties.subnet_prefix": {
          "value": $subnet_prefix
       }
    }')

echo "Configuring VMware NSX-T..."
om-linux \
  --target https://$OPSMAN_DOMAIN_OR_IP_ADDRESS \
  --skip-ssl-validation \
  --username "$OPS_MGR_USR" \
  --password "$OPS_MGR_PWD" \
  configure-product \
  --product-name VMware-NSX-T \
  --product-properties "${properties}"

