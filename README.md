# go-book-management
![maxresdefault](https://github.com/pooulad/go-book-management/assets/86445458/45e00ac9-69dd-4835-bd38-4e2d6015b496)

## API Reference

#### Get all Books

```http
  GET /book/
```
#### Create Book

```http
  POST /book/
```
```json
{
	"name": "mahdi",
	"author": "pooulad",
	"publication": "31/10//2023"
}
```

#### Get Book

```http
  GET /book/{bookId}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### Delete Book

```http
  DELETE /book/{bookId}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |



#### Get item

```http
  POST /book/{bookId}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |




