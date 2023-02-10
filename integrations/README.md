# Integrations

```bash
# DEV Env
 gcloud functions deploy user-signup --source . --trigger-event 'providers/firebase.auth/eventTypes/user.create' --trigger-resource 'contactkarma-dev' --runtime 'go116' --region 'us-central1' --set-secrets 'AC_API_KEY=projects/597119866017/secrets/activecampaign-api-key:1' --entry-point 'UserSignup'
```
