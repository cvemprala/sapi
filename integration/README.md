## Integration Tests

To run the integration tests, you need to have Python installed on your system. Follow the steps below:

1. Navigate to the root directory of the project.
2. Run the following command to execute all the integration tests:

```sh
python -m unittest discover -s integration
```

This command will discover and run all the integration tests in the `integration` folder.

To run a single integration test, use the following command:

```sh
python -m unittest integration.test_endpoints.TestEndpoints.test_create_todo
```

Replace `test_create_todo` with the name of the specific test method you want to run.
