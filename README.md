**Get Product**
----
  Retrieves a single product by a specified id.

* **URL:**
  `/product/{id}`

* **Method:**
  `GET`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:**
```json
{
    "id": 1,
    "name": "Example name",
    "price": 9.99,
    "currency": "GBP"
}
```

* **Error Resonse:**

  * **Code:** 500 <br />
    **Content:** `{ "message": "Error: Failed to retrieve product" }`

* **Sample Call:**

```bash
curl -X GET https://test.clerenet.com/product/1
```

**Get All Products**
----
  Retrieves all stored products.

* **URL:**
  `/product`

* **Method:**
  `GET`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:**
```json
[
    { "id": 1, "name": "Example name 1", "price": 2.99, "currency": "GBP"},
    { "id": 2, "name": "Example name 2", "price": 4.99, "currency": "GBP"},
    { "id": 3, "name": "Example name 3", "price": 9.99, "currency": "GBP"}
]
```

* **Error Resonse:**

  * **Code:** 500 <br />
    **Content:** `{ "message": "Error: Failed to retrieve products" }`

* **Sample Call:**

```bash
curl -X GET https://test.clerenet.com/product
```

**Create Product**
----
  Creates a product.

* **URL:**
  `/product`

* **Method:**
  `POST`

* **Data Params:**
```json
{
    "name": "string",
    "price": 0,
    "currency": "string"
}
```

* **Success Response:**

  * **Code:** 200 <br />
    **Content:**
```json
{
    "id": 1,
    "name": "Example name 1",
    "price": 2.99,
    "currency": "GBP"
}
```

* **Error Resonse:**

  * **Code:** 500 <br />
    **Content:** `{ "message": "Error: Failed to create product" }`

* **Sample Call:**

```bash
curl -X POST -d '{"name":"Test","price":10.99,"currency":"USD"}' https://test.clerenet.com/product
```

**Update Product**
----
  Updates a product, it uses the provided id in the data params to determine which product to update.

* **URL:**
  `/product`

* **Method:**
  `PUT`

* **Data Params:**
```json
{
    "id": 0,
    "name": "string",
    "price": 0,
    "currency": "string"
}
```

* **Success Response:**

  * **Code:** 200 <br />
    **Content:**
```json
{
    "id": 1,
    "name": "Updated name 1",
    "price": 2.99,
    "currency": "GBP"
}
```

* **Error Resonse:**

  * **Code:** 500 <br />
    **Content:** `{ "message": "Error: Failed to update product" }`

* **Sample Call:**

```bash
curl -X PUT -d '{"id": 1,"name":"Test","price":10.99,"currency":"USD"}' https://test.clerenet.com/product
```

**Delete Product**
----
  Deletes a product by the specified id.

* **URL:**
  `/product/{id}`

* **Method:**
  `DELETE`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `empty`

* **Error Resonse:**

  * **Code:** 500 <br />
    **Content:** `{ "message": "Error: Failed to delete product" }`

* **Sample Call:**

```bash
curl -X DELETE https://test.clerenet.com/product
```
