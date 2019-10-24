# Terraform Provider: Snowflake

----

**Please note**: If you believe you have found a security issue, _please responsibly disclose_ by contacting us at [security@chanzuckerberg.com](mailto:security@chanzuckerberg.com).

----

[![Join the chat at https://gitter.im/chanzuckerberg/terraform-provider-snowflake](https://badges.gitter.im/chanzuckerberg/terraform-provider-snowflake.svg)](https://gitter.im/chanzuckerberg/terraform-provider-snowflake?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge) [![Build Status](https://travis-ci.com/chanzuckerberg/terraform-provider-snowflake.svg?branch=master)](https://travis-ci.com/chanzuckerberg/terraform-provider-snowflake) [![codecov](https://codecov.io/gh/chanzuckerberg/terraform-provider-snowflake/branch/master/graph/badge.svg)](https://codecov.io/gh/chanzuckerberg/terraform-provider-snowflake)

This is a terraform provider plugin for managing [Snowflake](http://snowflakedb.com) accounts.

## Install

The easiest way is to run this command:

```shell
curl https://raw.githubusercontent.com/chanzuckerberg/terraform-provider-snowflake/master/download.sh | bash -s -- -b $HOME/.terraform.d/plugins
```

It runs a script generated by [godownloader](https://github.com/goreleaser/godownloader) which installs into the proper directory for terraform (~/.terraform.d/plugins).

You can also just download a binary from our [releases](https://github.com/chanzuckerberg/terraform-provider-snowflake/releases) and follow the [Terraform directions for installing 3rd party plugins](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins).

TODO fogg config

## Authentication

We currently support username + password and keypair auth and suggest that you do so via environment variables. Define a config with something like-

```hcl
provider "snowflake" {
  account = "..."
  role    = "..."
  region  = "..."
}
```

### Keypair Authentication Environment Variables
You should generate the public and private keys and set up environment variables. 

```shell
cd ~/.ssh
openssl genrsa -out snowflake_key 2048
openssl rsa -in snowflake_key -pubout -out snowflake_key.pub
```

To export the variables into your provider:
```shell
export SNOWFLAKE_USER="..."
export SNOWFLAKE_PRIVATE_KEY_PATH="~/.ssh/snowflake_key"
```

### Username and Password Environment Variables
If you choose to use Username and Password Authentication, export these credentials:
```shell
export SNOWFLAKE_USER='...'
export SNOWFLAKE_PASSWORD='...'
```


## Resources

We support managing a subset of snowflakedb resources, with a focus on access control and management.

You can see a number of examples [here](examples).

<!-- START -->

### snowflake_database

#### properties

|            NAME             |  TYPE  |                                  DESCRIPTION                                  | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|-----------------------------|--------|-------------------------------------------------------------------------------|----------|-----------|----------|---------|
| comment                     | string |                                                                               | true     | false     | false    | ""      |
| data_retention_time_in_days | int    |                                                                               | true     | false     | true     | <nil>   |
| from_share                  | map    | Specify a provider and a share in this map to create a database from a share. | true     | false     | false    | <nil>   |
| name                        | string |                                                                               | false    | true      | false    | <nil>   |

### snowflake_database_grant

Each grant resource is unique

#### properties

|     NAME      |  TYPE  |                      DESCRIPTION                       | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|---------------|--------|--------------------------------------------------------|----------|-----------|----------|---------|
| database_name | string | The name of the database on which to grant privileges. | false    | true      | false    | <nil>   |
| privilege     | string | The privilege to grant on the database.                | true     | false     | false    | "USAGE" |
| roles         | set    | Grants privilege to these roles.                       | true     | false     | false    | <nil>   |
| shares        | set    | Grants privilege to these shares.                      | true     | false     | false    | <nil>   |

### snowflake_managed_account

#### properties

|      NAME      |  TYPE  |                                                                  DESCRIPTION                                                                   | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT  |
|----------------|--------|------------------------------------------------------------------------------------------------------------------------------------------------|----------|-----------|----------|----------|
| admin_name     | string | Identifier, as well as login name, for the initial user in the managed account. This user serves as the account administrator for the account. | false    | true      | false    | <nil>    |
| admin_password | string | Password for the initial user in the managed account.                                                                                          | false    | true      | false    | <nil>    |
| cloud          | string | Cloud in which the managed account is located.                                                                                                 | false    | false     | true     | <nil>    |
| comment        | string | Specifies a comment for the managed account.                                                                                                   | true     | false     | false    | <nil>    |
| created_on     | string | Date and time when the managed account was created.                                                                                            | false    | false     | true     | <nil>    |
| locator        | string | Display name of the managed account.                                                                                                           | false    | false     | true     | <nil>    |
| name           | string | Identifier for the managed account; must be unique for your account.                                                                           | false    | true      | false    | <nil>    |
| region         | string | Snowflake Region in which the managed account is located.                                                                                      | false    | false     | true     | <nil>    |
| type           | string | Specifies the type of managed account.                                                                                                         | true     | false     | false    | "READER" |
| url            | string | URL for accessing the managed account, particularly through the web interface.                                                                 | false    | false     | true     | <nil>    |

### snowflake_resource_monitor

#### properties

|            NAME            |  TYPE  |                                                                   DESCRIPTION                                                                   | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|----------------------------|--------|-------------------------------------------------------------------------------------------------------------------------------------------------|----------|-----------|----------|---------|
| credit_quota               | int    | The number of credits allocated monthly to the resource monitor.                                                                                | true     | false     | true     | <nil>   |
| end_timestamp              | string | The date and time when the resource monitor suspends the assigned warehouses.                                                                   | true     | false     | false    | <nil>   |
| frequency                  | string | The frequency interval at which the credit usage resets to 0. If you set a frequency for a resource monitor, you must also set START_TIMESTAMP. | true     | false     | true     | <nil>   |
| name                       | string | Identifier for the resource monitor; must be unique for your account.                                                                           | false    | true      | false    | <nil>   |
| notify_triggers            | set    | A list of percentage thresholds at which to send an alert to subscribed users.                                                                  | true     | false     | false    | <nil>   |
| start_timestamp            | string | The date and time when the resource monitor starts monitoring credit usage for the assigned warehouses.                                         | true     | false     | true     | <nil>   |
| suspend_immediate_triggers | set    | A list of percentage thresholds at which to immediately suspend all warehouses.                                                                 | true     | false     | false    | <nil>   |
| suspend_triggers           | set    | A list of percentage thresholds at which to suspend all warehouses.                                                                             | true     | false     | false    | <nil>   |

### snowflake_role

#### properties

|  NAME   |  TYPE  | DESCRIPTION | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|---------|--------|-------------|----------|-----------|----------|---------|
| comment | string |             | true     | false     | false    | <nil>   |
| name    | string |             | false    | true      | false    | <nil>   |

### snowflake_role_grants

Each grant resource is unique

#### properties

|   NAME    |  TYPE  |              DESCRIPTION              | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|-----------|--------|---------------------------------------|----------|-----------|----------|---------|
| role_name | string | The name of the role we are granting. | false    | true      | false    | <nil>   |
| roles     | set    | Grants role to this specified role.   | true     | false     | false    | <nil>   |
| users     | set    | Grants role to this specified user.   | true     | false     | false    | <nil>   |

### snowflake_schema

#### properties

|        NAME         |  TYPE  |                                                                                                                      DESCRIPTION                                                                                                                       | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|---------------------|--------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|-----------|----------|---------|
| comment             | string | Specifies a comment for the schema.                                                                                                                                                                                                                    | true     | false     | false    | <nil>   |
| data_retention_days | int    | Specifies the number of days for which Time Travel actions (CLONE and UNDROP) can be performed on the schema, as well as specifying the default Time Travel retention time for all tables created in the schema.                                       | true     | false     | false    |       1 |
| database            | string | The database in which to create the schema.                                                                                                                                                                                                            | false    | true      | false    | <nil>   |
| is_managed          | bool   | Specifies a managed schema. Managed access schemas centralize privilege management with the schema owner.                                                                                                                                              | true     | false     | false    | false   |
| is_transient        | bool   | Specifies a schema as transient. Transient schemas do not have a Fail-safe period so they do not incur additional storage costs once they leave Time Travel; however, this means they are also not protected by Fail-safe in the event of a data loss. | true     | false     | false    | false   |
| name                | string | Specifies the identifier for the schema; must be unique for the database in which the schema is created.                                                                                                                                               | false    | true      | false    | <nil>   |

### snowflake_schema_grant

Each grant resource is unique

#### properties

|     NAME      |  TYPE  |                                                                  DESCRIPTION                                                                  | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|---------------|--------|-----------------------------------------------------------------------------------------------------------------------------------------------|----------|-----------|----------|---------|
| database_name | string | The name of the database containing the schema on which to grant privileges.                                                                  | false    | true      | false    | <nil>   |
| privilege     | string | The privilege to grant on the schema.  Note that if "OWNERSHIP" is specified, ensure that the role that terraform is using is granted access. | true     | false     | false    | "USAGE" |
| roles         | set    | Grants privilege to these roles.                                                                                                              | true     | false     | false    | <nil>   |
| schema_name   | string | The name of the schema on which to grant privileges.                                                                                          | false    | true      | false    | <nil>   |
| shares        | set    | Grants privilege to these shares.                                                                                                             | true     | false     | false    | <nil>   |

### snowflake_share

#### properties

|   NAME   |  TYPE  |                                              DESCRIPTION                                              | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|----------|--------|-------------------------------------------------------------------------------------------------------|----------|-----------|----------|---------|
| accounts | set    | A list of accounts to be added to the share.                                                          | true     | false     | false    | <nil>   |
| comment  | string | Specifies a comment for the managed account.                                                          | true     | false     | false    | <nil>   |
| name     | string | Specifies the identifier for the share; must be unique for the account in which the share is created. | false    | true      | false    | <nil>   |

### snowflake_table_grant

Each grant resource is unique

#### properties

|     NAME      |  TYPE  |                                                                           DESCRIPTION                                                                           | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT  |
|---------------|--------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|-----------|----------|----------|
| database_name | string | The name of the database containing the current or future tables on which to grant privileges.                                                                  | false    | true      | false    | <nil>    |
| on_future     | bool   | When this is set to true, apply this grant on all future tables in the given schema.  The table_name and shares fields must be unset in order to use on_future. | true     | false     | false    | false    |
| privilege     | string | The privilege to grant on the current or future table.                                                                                                          | true     | false     | false    | "SELECT" |
| roles         | set    | Grants privilege to these roles.                                                                                                                                | true     | false     | false    | <nil>    |
| schema_name   | string | The name of the schema containing the current or future tables on which to grant privileges.                                                                    | true     | false     | false    | "PUBLIC" |
| shares        | set    | Grants privilege to these shares (only valid if on_future is unset).                                                                                            | true     | false     | false    | <nil>    |
| table_name    | string | The name of the table on which to grant privileges immediately (only valid if on_future is unset).                                                              | true     | false     | false    | <nil>    |

### snowflake_user

#### properties

|         NAME         |  TYPE  |                                                                                                        DESCRIPTION                                                                                                         | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|----------------------|--------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|-----------|----------|---------|
| comment              | string |                                                                                                                                                                                                                            | true     | false     | false    | <nil>   |
| default_namespace    | string | Specifies the namespace (database only or database and schema) that is active by default for the user’s session upon login.                                                                                                | true     | false     | false    | <nil>   |
| default_role         | string | Specifies the role that is active by default for the user’s session upon login.                                                                                                                                            | true     | false     | true     | <nil>   |
| default_warehouse    | string | Specifies the virtual warehouse that is active by default for the user’s session upon login.                                                                                                                               | true     | false     | false    | <nil>   |
| disabled             | bool   |                                                                                                                                                                                                                            | true     | false     | true     | <nil>   |
| has_rsa_public_key   | bool   | Will be true if user as an RSA key set.                                                                                                                                                                                    | false    | false     | true     | <nil>   |
| login_name           | string | The name users use to log in. If not supplied, snowflake will use name instead.                                                                                                                                            | true     | false     | true     | <nil>   |
| must_change_password | bool   | Specifies whether the user is forced to change their password on next login (including their first/initial login) into the system.                                                                                         | true     | false     | false    | <nil>   |
| name                 | string | Name of the user. Note that if you do not supply login_name this will be used as login_name. [doc](https://docs.snowflake.net/manuals/sql-reference/sql/create-user.html#required-parameters)                              | false    | true      | false    | <nil>   |
| password             | string | **WARNING:** this will put the password in the terraform state file. Use carefully.                                                                                                                                        | true     | false     | false    | <nil>   |
| rsa_public_key       | string | Specifies the user’s RSA public key; used for key-pair authentication. Must be on 1 line without header and trailer.                                                                                                       | true     | false     | false    | <nil>   |
| rsa_public_key_2     | string | Specifies the user’s second RSA public key; used to rotate the public and private keys for key-pair authentication based on an expiration schedule set by your organization. Must be on 1 line without header and trailer. | true     | false     | false    | <nil>   |

### snowflake_view

#### properties

|   NAME    |  TYPE  |                                                          DESCRIPTION                                                          | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT  |
|-----------|--------|-------------------------------------------------------------------------------------------------------------------------------|----------|-----------|----------|----------|
| comment   | string | Specifies a comment for the view.                                                                                             | true     | false     | false    | <nil>    |
| database  | string | The database in which to create the view. Don't use the | character.                                                          | false    | true      | false    | <nil>    |
| is_secure | bool   | Specifies that the view is secure.                                                                                            | true     | false     | false    | false    |
| name      | string | Specifies the identifier for the view; must be unique for the schema in which the view is created. Don't use the | character. | false    | true      | false    | <nil>    |
| schema    | string | The schema in which to create the view. Don't use the | character.                                                            | true     | false     | false    | "PUBLIC" |
| statement | string | Specifies the query used to create the view.                                                                                  | false    | true      | false    | <nil>    |

### snowflake_view_grant

Each grant resource is unique

#### properties

|     NAME      |  TYPE  |                                                                          DESCRIPTION                                                                          | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT  |
|---------------|--------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|-----------|----------|----------|
| database_name | string | The name of the database containing the current or future views on which to grant privileges.                                                                 | false    | true      | false    | <nil>    |
| on_future     | bool   | When this is set to true, apply this grant on all future views in the given schema.  The view_name and shares fields must be unset in order to use on_future. | true     | false     | false    | false    |
| privilege     | string | The privilege to grant on the current or future view.                                                                                                         | true     | false     | false    | "SELECT" |
| roles         | set    | Grants privilege to these roles.                                                                                                                              | true     | false     | false    | <nil>    |
| schema_name   | string | The name of the schema containing the current or future views on which to grant privileges.                                                                   | true     | false     | false    | "PUBLIC" |
| shares        | set    | Grants privilege to these shares (only valid if on_future is unset).                                                                                          | true     | false     | false    | <nil>    |
| view_name     | string | The name of the view on which to grant privileges immediately (only valid if on_future is unset).                                                             | true     | false     | false    | <nil>    |

### snowflake_warehouse

#### properties

|         NAME          |  TYPE  |                                                               DESCRIPTION                                                                | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|-----------------------|--------|------------------------------------------------------------------------------------------------------------------------------------------|----------|-----------|----------|---------|
| auto_resume           | bool   | Specifies whether to automatically resume a warehouse when a SQL statement (e.g. query) is submitted to it.                              | true     | false     | true     | <nil>   |
| auto_suspend          | int    | Specifies the number of seconds of inactivity after which a warehouse is automatically suspended.                                        | true     | false     | true     | <nil>   |
| comment               | string |                                                                                                                                          | true     | false     | false    | ""      |
| initially_suspended   | bool   | Specifies whether the warehouse is created initially in the ‘Suspended’ state.                                                           | true     | false     | false    | <nil>   |
| max_cluster_count     | int    | Specifies the maximum number of server clusters for the warehouse.                                                                       | true     | false     | true     | <nil>   |
| min_cluster_count     | int    | Specifies the minimum number of server clusters for the warehouse (only applies to multi-cluster warehouses).                            | true     | false     | true     | <nil>   |
| name                  | string |                                                                                                                                          | false    | true      | false    | <nil>   |
| resource_monitor      | string | Specifies the name of a resource monitor that is explicitly assigned to the warehouse.                                                   | true     | false     | true     | <nil>   |
| scaling_policy        | string | Specifies the policy for automatically starting and shutting down clusters in a multi-cluster warehouse running in Auto-scale mode.      | true     | false     | true     | <nil>   |
| wait_for_provisioning | bool   | Specifies whether the warehouse, after being resized, waits for all the servers to provision before executing any queued or new queries. | true     | false     | false    | <nil>   |
| warehouse_size        | string |                                                                                                                                          | true     | false     | true     | <nil>   |

### snowflake_warehouse_grant

Each grant resource is unique

#### properties

|      NAME      |  TYPE  |                       DESCRIPTION                       | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|----------------|--------|---------------------------------------------------------|----------|-----------|----------|---------|
| privilege      | string | The privilege to grant on the warehouse.                | true     | false     | false    | "USAGE" |
| roles          | set    | Grants privilege to these roles.                        | true     | false     | false    | <nil>   |
| warehouse_name | string | The name of the warehouse on which to grant privileges. | false    | true      | false    | <nil>   |
<!-- END -->

## Development

To do development you need Go installed, this repo cloned and that's about it. It has not been tested on Windows, so if you find problems let us know.

If you want to build and test the provider localling there is a make target `make install-tf` that will build the provider binary and install it in a location that terraform can find.

### Testing

For the Terraform resources, there are 3 levels of testing - internal, unit and acceptance tests.

The 'internal' tests are run in the `github.com/chanzuckerberg/terraform-provider-snowflake/pkg/resources` package so that they can test functions that are not exported. These tests are intended to be limited to unit tests for simple functions.

The 'unit' tests are run in  `github.com/chanzuckerberg/terraform-provider-snowflake/pkg/resources_test`, so they only have access to the exported methods of `resources`. These tests exercise the CRUD methods that on the terraform resources. Note that all tests here make use of database mocking and are run locally. This means the tests are fast, but are liable to be wrong in suble ways (since the mocks are unlikely to be perfect).

You can run these first two sets of tests with `make test`.

The 'acceptance' tests run the full stack, creating, modifying and destroying resources in a live snowflake account. To run them you need a snowflake account and the proper environment variables set- SNOWFLAKE_ACCOUNT, SNOWFLAKE_USER, SNOWFLAKE_PASSWORD, SNOWFLAKE_ROLE. These tests are slower but have higher fidelity.

To run all tests, including the acceptance tests, run `make test-acceptance`.

Note that we also run all tests in our [Travis-CI account](https://travis-ci.com/chanzuckerberg/terraform-provider-snowflake).
