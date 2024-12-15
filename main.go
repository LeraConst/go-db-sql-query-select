package main

import (
	"fmt"
	"database/sql"

	_ "modernc.org/sqlite"
)

type Sale struct {
	Product int
	Volume  int
	Date    string
}

// String реализует метод интерфейса fmt.Stringer для Sale, возвращает строковое представление объекта Sale.
// Теперь, если передать объект Sale в fmt.Println(), то выведется строка, которую вернёт эта функция.
func (s Sale) String() string {
	return fmt.Sprintf("Product: %d Volume: %d Date:%s", s.Product, s.Volume, s.Date)
}

func selectSales(client int) ([]Sale, error) {
	var sales []Sale

	 // Открываем соединение с базой данных SQLite
	 db, err := sql.Open("sqlite", "demo.db")
	 if err != nil {
		 return nil, err
	 }
	 defer db.Close() // Закрываем соединение с базой данных по завершении работы


	 // Выполняем SQL-запрос для получения данных
	 row, err := db.Query("SELECT product, volume, date FROM sales WHERE client = :client", sql.Named("client", client))
	 if err != nil {
		return nil, err
	 }
	 defer row.Close() // Закрываем результат запроса по завершении работы


	  // Проходим по всем строкам результата запроса
	  for row.Next() {
		var product, volume int 
		var date string
        // Считываем данные из текущей строки в поля структуры Product
        err := row.Scan(&product, &volume, &date)
        if err != nil {
            return nil, err
        }
		// Создаём объект Sale и добавляем его в срез
		sale := Sale{
			Product: product,
			Volume:  volume,
			Date:    date,
		}
		sales = append(sales, sale)
	}

	// Проверяем наличие ошибок при итерации
	if err := row.Err(); err != nil {
		return nil, err
    }
	return sales, nil
}

func main() {
	client := 208

	sales, err := selectSales(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sale := range sales {
		fmt.Println(sale)
	}
}
