# DATA TYPES

## **`mdstrings`**

```go
type MdString string
```

> question and mcq strings are markdown encoded (need to check for xxs and script injection)

## **`ID`:**

```go
type ID string
```

> IDs are string encoded twitter-like snowflakes (41 bit for timestamp; 10 for node; 12 bits for step)

## **`MCQ`:**

```go
type MCQ struct {
    ID snowflakes.ID
    LastMod int64
    Title string
    Question MdString
    Answers []MdString
    correct []int
    Explanation MdString
    Author snowflakes.ID
    Ref []string
    LastRef int64
}
```

```json
{
    "id":"57832695372832",
    "last-mod": 6947637864830,
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
    "author":"7054704732965",
    "_ref":["s1-math"],
    "_last-ref" : 785478365873408,
}
```

> _ref is used as a GC; to keep track to which repo it is added (when removed everywhere don't know if to erase it or keep it, maybe keep it for 30 days or smth.)

---

## **`question`:**

```go
type Question struct {
    ID snowflakes.ID
    LastMod int64
    Name string
    Question MdString
    Explanation MdString
    Author snowflakes.ID
    Ref []string
    LastRef int64
}
```

```json
{
    "id":"40863296589",
    "last-mod": 6947637864830,
    "title":"incredibly hard math question",
    "question":"what's after 4",
    "explanation":"according to the lesson: 1, 2, 3, 4, 5, 6, 7, 9;
    5 is after 4 so the answer is 5",
    "author":"105634786578",
    "_ref": ["s1-math"],
    "_last-ref" : 785478365873408,
}
```

## **`Category`:**

```go
type Category struct {
    ID snowflakes.ID
    Name string
    Mcqs []MCQ
    Questions []Question
}
```

```json
{
    "id":"19075043703290",
    "name":"math-s1",
    "mcq-ids":[
        "57832695372832",
    ],
    "question-ids":[
        "40863296589",
    ]
}
```
