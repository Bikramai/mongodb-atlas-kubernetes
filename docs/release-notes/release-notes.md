
# MongoDB Atlas Operator v0.8.0

## Atlas Operator

* Changes
    *  Upgrade Controller Runtime to v0.11.0
    *  Upgraded to Go 1.17
    *  When installing a cluster using the helm chart, helm will not exit until the cluster is ready if `postInstallHook.enabled` is set to true.
    *  The operator now only watches secrets with the label `atlas.mongodb.com/type=credentials` to avoid watching unnecessary secrets.
    *  It is possible to configure the Operator to skip reconciliations on specific resources by adding the annotation `mongodb.com/atlas-reconciliation-policy=skip`.
    * Enable User Authentication using X.509 Certificates
* Bug fixes 
    * Fixed an issue where errors would be logged upon resource deletion.


## AtlasProject Resource
* Changes
    * The AtlasProject will not be marked as ready until the Project IP Access List is successfully created.
  
*The images can be found in:*

https://quay.io/repository/mongodb/mongodb-atlas-operator