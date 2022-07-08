https://github.com/GoogleCloudPlatform/cloud-builders-community/tree/master/ko

```bash
git clone git@github.com:GoogleCloudPlatform/cloud-builders-community.git
cd cloud-builders-community/ko
gcloud auth login
gcloud config set project ${GCP_PROJECT}
gcloud builds submit . --config=cloudbuild.yaml --substitutions=_KO_GIT_TAG="v0.10.0"
```

バージョンをv0.10.0より上げると動かない・・・。