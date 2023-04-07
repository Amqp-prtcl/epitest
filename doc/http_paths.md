# HTTP PATHS

## HTTP codes

- **HTTP500** -> connection problem with the MongoDB Instance
- **HTTP400** -> malformed or invalid request
- **HTTP404** -> id or category name not found
- **HTTP200** -> request succeeded with non-empty body
- **HTTP201** -> object successfully created
- **HTTP204** -> request succeeded with empty body

## Getters

> Are done with the **GET** Every time an ID is asked as a parameter **only in getter requests**, `bulk` can be specified instead; in this case the request body must contain a valid json encoded list of IDs and the server will respond with an array of every id that matched or in the case of category added or removed,
(Note that the max number of ids per request may be capped)

## Modifiers

> Are done with **PUT** http verb and the `id` url parameter. `bulk` request are not allowed; to modify a card or question, you'll need to add the same body as for creation but with only the fields you want to modify (Note: any field found in request body that is not part of the mutable fields of the object will be discarded)

### Example

- url **PUT** `/api/mcqs?id=56473289478`

- request body:

```json
{
  "title":"incredibly hard prog mcq",
}
```

- server response: HTTP204

---

## **MCQs**

### **POST** `/api/mcq/new/`

#### Expected bedy

  ```json
  {
   "title":"addition: 1",
    "question":"1+1 ?",
    "answers": [
        "2",
        "3",
        "1",
        "5",
    ],
    "correct":[
        0,
    ],
    "explanation":"I + I = II (2)",
  }
  ```

  following struct

  ```go
  struct {
    Title string
    Question MdString
    Answers []MdString
    Correct []int
    Explanation MdString
  }
  ```

### **GET** `/api/mcqs/`

see **GET** `api/questions/`

### **PUT** `/api/mcqs/`

see **PUT** `api/questions/`

---

## **Questions**

### **POST** `/api/question/new/`

#### Expected body

  ```json
  {
    "title":"incredibly hard math question",
    "question":"what's after 4",
    "explanation":"according to the lesson: 1, 2, 3, 4, 5, 6, 7, 9;
    5 is after 4 so the answer is 5",
  }
  ```

  following struct:

  ```go
  struct {
    Title string
    Question MdString
    Explanation MdString
  }
  ```

### **GET** `/api/questions/`

#### URL parameters (in order of priority)

- `id`: takes an id and returns corresponding question (or Http404 if it is an invalid ID)

#### Examples

- id request:

  - url : **GET** `/api/questions?id=67856378943`

  - server response:

    ```json
    {
        "id":"40863296589",
        "title":"incredibly hard math question",
        "question":"what's after 4",
        "explanation":"according to the lesson: 1, 2, 3, 4, 5, 6, 7, 9;
        5 is after 4 so the answer is 5",
    }
    ```

- bulk request:

  - url: **GET** `/api/questions?id=bulk`

  - request body:

    ```json
    [
        "78564783234",
        "7580439659743",
        "678564378985493",
        "789543689674",
    ]
    ```

  - server response:

    ```json
    [
        {
            "id":"40863296589",
            "title":"incredibly hard math question",
            "question":"what's after 4",
            "explanation":"according to the lesson: 1, 2, 3, 4, 5, 6, 7, 9;
            5 is after 4 so the answer is 5",
        },
        {
            "id":"40863296589",
            "title":"incredibly hard math question",
            "question":"what's after 4",
            "explanation":"according to the lesson: 1, 2, 3, 4, 5, 6, 7, 9;
            5 is after 4 so the answer is 5",
        },
        {
            "id":"40863296589",
            "title":"incredibly hard math question",
            "question":"what's after 4",
            "explanation":"according to the lesson: 1, 2, 3, 4, 5, 6, 7, 9;
            5 is after 4 so the answer is 5",
        },
    ]
    ```

### **PUT** `/api/question/`

---

## **Categories**

### **POST** `/api/category/new?name=<custom-name>`

> Creates an empty category with given name (or Http409 if a category already exists with that name)

#### URL parameters

- `name`: specifies the category name (Mandatory, not using this parameter will result in a HTTP400 from server)

### **GET** `/api/categories/`

returns a list of all existing categories

response example:

```json
[
    {
        "name":"math-s1",
        "id":"689546378932",
    },
]
```

### **GET** `/api/category/{category-name}/`

if a category with corresponding name exists, server responds with a json encoded category data types

### **PUT** `/api/category/{category-name}/`

#### URL parameters (in order or priority)

- `add`: takes an id and adds it to the cat if not already present

  > If used as a bulk request, the server will return a json list of ids that were not able to be added to category so an empty list means a full success (a bulk rm request will just ignore ids that are not already inside category)

- `rm`: takes an id and removes it if in category

  > if removed question or mcq is only present in this category, it won't be accessible without its ID and will be remove after 30 days (within this time limit, you can still retrieve it by adding it to a category).

- `name`: takes a string and edits the name of the category ( or Http409 if a category already exists with provided name )

#### Examples

- name request

  - url: **PUT** `api/category/s1-ma?name=s1-math`

  - server response: HTTP204

- remove request:

  - url: **PUT** `api/category/s1-math?rm=678546398934`

  - server response: HTTP204

- add request:

  - url: **PUT** `api/category/s1-math?add=6574839675843`

  - server response: HTTP204

- bulk request:

  - url: **PUT** `api/category.s1-math?add=bulk`

  - request body:

    ```json
    [
      "6754839786783",
      "4302",
      "1975436212121",
      "1112378943927",
      "1118433784253",
    ]
    ```

  - server response:

    ```json
    [
      "4302",
      "1975436212121",
      "1118433784253",
    ]
    ```

---
