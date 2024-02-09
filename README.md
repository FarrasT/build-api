# build-api

## Alur Kerja
1. API ini digunakan untuk aplikasi pencatatan transaksi dagangan offline
2. Pengguna aplikasi adalah pedagang
3. Pedagang bisa mencatat transaksi nya secara langsung pada aplikasi ini atau mendigitalisasi catatan mereka yang sudah lama
4. Pedagang menyimpan catatannya dengan API `Create Order`. Tersimpan `customer_name`, `order_at` dan `items`
5. Pedagang bisa mengambil semua catatan transaksinya dengan api `Get All Data`
6. Pedagang bisa menghapus `order` menggunakan API `Delete Order`, maka akan terhapus juga semua `items` milik `order` tersebut. 
7. Pedagang bisa menghapus salah satu `items` saja, tidak menghapus semua `items` ataupun `order` nya. Menggunakan API `Delete Items`
8. Pedagang dapat mengubah data `order` dan `items` dengan API `Update Order`. API ini memiliki beberapa ketentuan berikut ini:
* Apabila kombinasi `item_code` dan `id` pada database dan body request ada. Maka data `items` di database akan berubah sesuai dengan request
* Apabila kombinasi `item_code` dan `id` di request body ternyata tidak ada di database. Maka `items` baru akan tersimpan
* Apabila ada 2 `item_code` sama pada request body, maka data yang tersimpan adalah data `item_code` tersebut dengan index terbesar pada array `items` di request body


## API
### Create Order
```
POST /api/order HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Content-Length: 281

{
	"customer_name" : "Farras Timorremboko",
	"order_at" : "09-02-2024T10:00:00", 
	"items" : [
		{
			"item_code" : "123",
			"description" : "Iphone",
			"quantity" : 1
		},
        {
			"item_code" : "234",
			"description" : "Iphone X",
			"quantity" : 9
		}
	]
}
```
Response
```
{
    "code": 201,
    "data": {
        "customer_name": "Farras Timorremboko",
        "order_at": "09-02-2024T10:00:00",
        "items": [
            {
                "item_code": "123",
                "description": "Iphone",
                "quantity": 1
            },
            {
                "item_code": "234",
                "description": "Iphone X",
                "quantity": 1
            }
        ]
    },
    "message": "success"
}
```


### Get All Data
```
GET /api/order HTTP/1.1
Host: localhost:8080
```
Response
```
{
    "code": 200,
    "data": [
        {
            "id": 7,
            "customer_name": "Farras Tinjur",
            "items": [
                {
                    "item_id": 8,
                    "item_code": "123",
                    "order_id": 7,
                    "description": "Iphone",
                    "quantity": 1,
                    "created_at": "2024-02-09T11:02:10.493283+07:00",
                    "updated_at": "2024-02-09T11:02:10.493283+07:00"
                },
                {
                    "item_id": 11,
                    "item_code": "456",
                    "order_id": 7,
                    "description": "Iphone 18",
                    "quantity": 2,
                    "created_at": "2024-02-09T13:48:08.345684+07:00",
                    "updated_at": "2024-02-09T13:48:08.345684+07:00"
                }
            ],
            "created_at": "2024-02-09T11:02:10.487951+07:00",
            "updated_at": "2024-02-09T13:48:08.337187+07:00"
        }
    ],
    "message": "success"
}
```

### Delete Order
```
DELETE /api/order/:orderId HTTP/1.1
Host: localhost:8080
```
Response
```
{
    "code": 200,
    "message": "success"
}
```

### Delete Items
```
DELETE /api/order/:orderId/item/:item_id HTTP/1.1
Host: localhost:8080
```
Response
```
{
    "code": 200,
    "message": "success"
}
```

### Update Order
```
PUT /api/order/:orderId HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Content-Length: 354

{
    "customer_name": "Farras Tinjur",
    "order_at": "09-02-2024T10:00:00",
    "items": [
        {
            "item_code": "789",
            "description": "Iphone 16",
            "quantity": 8
        },
        {
            "item_code": "789",
            "description": "Iphone 23",
            "quantity": 11
        }
    ]
}
```
Response
```
{
    "code": 200,
    "data": {
        "customer_name": "Farras Timor",
        "order_at": "09-02-2024T10:00:00",
        "items": [
            {
                "item_code": "789",
                "description": "Iphone 16",
                "quantity": 8
            },
            {
                "item_code": "789",
                "description": "Iphone 23",
                "quantity": 11
            }
        ]
    },
    "message": "success"
}
```