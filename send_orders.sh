#!/bin/bash

FILE="orders.json"
URL="http://localhost:8080/api/v1/orders"

if [ ! -f "$FILE" ]; then
  echo "Файл $FILE не найден!"
  exit 1
fi

COUNT=$(jq length "$FILE")

echo "Параллельная отправка $COUNT заказов на $URL"

MAX_PARALLEL=10  # Максимум одновременно работающих процессов
JOBS=0

for ((i = 0; i < COUNT; i++)); do
  ORDER=$(jq -c ".[$i]" "$FILE")

  {
    echo "Отправка заказа $((i+1))..."
    curl -s -o /dev/null -w "HTTP статус: %{http_code}\n" -X POST "$URL" \
      -H "Content-Type: application/json" \
      -d "$ORDER"
  } &

  ((JOBS++))

  # Ограничиваем количество параллельных задач
  if (( JOBS >= MAX_PARALLEL )); then
    wait -n  # Ждём завершения хотя бы одного процесса
    ((JOBS--))
  fi
done

# Дожидаемся завершения всех оставшихся задач
wait

echo "Готово."
