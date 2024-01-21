package utils

import "reflect"

func FindByID[T any](items []T, targetID string) *T {
	for _, item := range items {
		// reflectを使用してIdフィールドにアクセスする
		value := reflect.ValueOf(item)
		idField := value.FieldByName("Id")

		if idField.IsValid() && idField.String() == targetID {
			return &item
		}
	}
	return nil
}
