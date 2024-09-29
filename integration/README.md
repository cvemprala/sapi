## Integration Tests

To run the integration tests, you need to have Python installed on your system. Follow the steps below:

1. Navigate to the root directory of the project.
2. Run the following command to execute all the integration tests:

```sh
pytest integration
```

This command will discover and run all the integration tests in the `integration` folder.

To run a single integration test, use the following command:

```sh
pytest integration/test_endpoints.py::TestEndpoints::test_create
```

Replace `test_create` with the name of the specific test method you want to run.
