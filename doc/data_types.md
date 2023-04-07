# DATA TYPES

## **`CkStrings`**

> see <https://ckeditor.com/docs/ckeditor5/latest/installation/advanced/saving-data.html>

```go
type CkString string
```

> question and mcq strings are markdown encoded (need to check for xxs and script injection)

## **`ID`:**

```go
type ID string
```

> IDs are string encoded twitter-like snowflakes (41 bit for timestamp; 10 for node; 12 bits for step)

## **`HEADER`**

```go
type Header struct {
    ID       snowflakes.ID `json:"id"`
    LastMod  int64         `json:"last-mod"`
    Title    string        `json:"title"`
    Category []string      `json:"categories"`
}
```

## **`MCQ`:**

```go
type MCQ struct {
    Header
    Question    CkString      `json:"question"`
    Answers     []CkString    `json:"answers"`
    Correct     []int         `json:"correct"`
    Explanation CkString      `json:"explanation"`
    Author      snowflakes.ID `json:"author"`
    Ref         []string      `json:"ref"`
    LastRef     int64         `json:"ref-mod"`
}
```

> _ref is used as a GC; to keep track to which repo it is added (when removed everywhere don't know if to erase it or keep it, maybe keep it for 30 days or smth.)

---

## **`question`:**

```go
type Question struct {
    Header
    Question    CkString      `json:"question"`
    Explanation CkString      `json:"explanation"`
    Author      snowflakes.ID `json:"author"`
    Ref         []string      `json:"ref"`
    LastRef     int64         `json:"ref-mod"`
}
```

## **`Category`:**

```go
type Group struct {
    ID        snowflakes.ID `json:"id"`
    Name      string        `json:"name"`
    Mcqs      []Header      `json:"mcqs"`
    Questions []Header      `json:"questions"`
}
```
