import requests
import unittest

class TestEndpoints(unittest.TestCase):

    def setUp(self):
        self.base_url = "http://localhost:8080/api/todos"

    def test_1(self):
        print("Testing create_todo")
        payload = {
            "Title": "Buy groceries",
            "Description": "Milk, Bread, Eggs",
            "DueDate": "2023-12-31T23:59:59Z",
            "Priority": 1,
            "Tags": ["shopping", "errands"]
        }
        response = requests.post(self.base_url, json=payload)
        self.assertEqual(response.status_code, 201)
        self.assertEqual(response.json()["Title"], "Buy groceries")

    def test_2(self):
        print("Testing get_todo")
        response = requests.get(f"{self.base_url}/1")
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json()["ID"], 1)

    def test_3(self):
        print("Testing update_todo")
        payload = {
            "Title": "Buy groceries and more",
            "Description": "Milk, Bread, Eggs, and Butter",
            "DueDate": "2024-01-01T23:59:59Z",
            "Priority": 2,
            "Tags": ["shopping", "errands", "important"]
        }
        response = requests.put(f"{self.base_url}/1", json=payload)
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json()["Title"], "Buy groceries and more")

    def test_4(self):
        print("Testing delete_todo")
        response = requests.delete(f"{self.base_url}/1")
        self.assertEqual(response.status_code, 204)

if __name__ == "__main__":
    unittest.main()
