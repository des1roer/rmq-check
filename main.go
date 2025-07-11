package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
	_ "rmq-check/internal/config"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	// 1. Конфигурация подключения
	endpoint := os.Getenv("ENDPOINT")
	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	useSSL := true

	// 2. Создание клиента MinIO
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		logrus.Fatalf("Ошибка создания клиента: %v", err)
	}

	// 3. Подготовка JSON данных
	data := map[string]interface{}{
		"name":    "Тестовый документ",
		"value":   42,
		"tags":    []string{"go", "minio"},
		"created": time.Now(),
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		logrus.Fatalf("Ошибка сериализации JSON: %v", err)
	}

	// 4. Параметры загрузки
	bucketName := "test-bucket"
	objectName := "test-data.json"
	contentType := "application/json"
	reader := bytes.NewReader(jsonData)
	objectSize := int64(reader.Len())

	// 5. Создание бакета (если не существует)
	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		logrus.Fatalf("Проверка бакета: %v", err)
	}

	if !exists {
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			logrus.Fatalf("Создание бакета: %v", err)
		}
		logrus.Printf("Бакет %s создан\n", bucketName)
	}

	// 6. Загрузка файла в MinIO
	uploadInfo, err := minioClient.PutObject(
		ctx,
		bucketName,
		objectName,
		reader,
		objectSize,
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		logrus.Fatalf("Ошибка загрузки: %v", err)
	}

	logrus.Printf("Файл успешно загружен: %v\n", uploadInfo)
}
