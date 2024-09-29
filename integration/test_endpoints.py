import requests
import unittest

class TestEndpoints(unittest.TestCase):

    def setUp(self):
        self.base_url = "http://localhost:8080/api/todos"

    def test_create(self):
        print("Testing create_todo")
        payload = {
            "title": "Buy groceries",
            "description": "Milk, Bread, Eggs",
            "dueDate": "2023-12-31T23:59:59Z",
            "priority": 1,
            "tags": ["shopping", "errands"]
        }
        response = requests.post(self.base_url, json=payload)
        self.assertEqual(response.status_code, 201)
        self.assertEqual(response.json()["Title"], "Buy groceries")

    def test_get(self):
        print("Testing get_todo")
        response = requests.get(f"{self.base_url}/1")
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json()["ID"], 1)

    def test_update(self):
        print("Testing update_todo")
        payload = {
            "title": "Buy groceries and more",
            "description": "Milk, Bread, Eggs, and Butter",
            "dueDate": "2024-01-01T23:59:59Z",
            "priority": 2,
            "tags": ["shopping", "errands", "important"]
        }
        response = requests.put(f"{self.base_url}/1", json=payload)
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json()["title"], "Buy groceries and more")

    def test_delete(self):
        print("Testing delete_todo")
        response = requests.delete(f"{self.base_url}/1")
        self.assertEqual(response.status_code, 204)

if __name__ == "__main__":
    unittest.main()
