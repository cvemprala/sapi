import requests
import pytest

@pytest.fixture
def base_url():
    return "http://localhost:8080/api/todos"

def test_create(base_url):
    print("Testing create_todo")
    payload = {
        "title": "Buy groceries",
        "description": "Milk, Bread, Eggs",
        "dueDate": "2023-12-31T23:59:59Z",
        "priority": 1,
        "tags": ["shopping", "errands"]
    }
    response = requests.post(base_url, json=payload)
    assert response.status_code == 201
    assert response.json()["Title"] == "Buy groceries"

def test_get(base_url):
    print("Testing get_todo")
    response = requests.get(f"{base_url}/1")
    assert response.status_code == 200
    assert response.json()["ID"] == 1

def test_update(base_url):
    print("Testing update_todo")
    payload = {
        "title": "Buy groceries and more",
        "description": "Milk, Bread, Eggs, and Butter",
        "dueDate": "2024-01-01T23:59:59Z",
        "priority": 2,
        "tags": ["shopping", "errands", "important"]
    }
    response = requests.put(f"{base_url}/1", json=payload)
    assert response.status_code == 200
    assert response.json()["title"] == "Buy groceries and more"

def test_delete(base_url):
    print("Testing delete_todo")
    response = requests.delete(f"{base_url}/1")
    assert response.status_code == 204
