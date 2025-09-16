#!/bin/bash
set -e

# Папка для сохранения данных
OUT_DIR="./data"

# Количество маршрутов для каждой категории
COUNT=100

# Общее количество точек в каждом маршруте
POINTS=200

# Начальные и конечные координаты
FROMX=0
FROMY=0
TOX=100
TOY=100

echo "🚀 Генерация маршрутов в $OUT_DIR ..."

# ---------- НОРМАЛЬНЫЕ ----------
echo "📌 Генерация нормальных маршрутов..."
go run . \
  -fromX $FROMX -fromY $FROMY -toX $TOX -toY $TOY \
  -points $POINTS -type normal -count $COUNT -out $OUT_DIR

# ---------- АНОМАЛИЯ: ZIGZAG ----------
echo "📌 Генерация маршрутов с аномалией: zigzag..."
go run . \
  -fromX $FROMX -fromY $FROMY -toX $TOX -toY $TOY \
  -points $POINTS -type abnormal -anom zigzag -count $COUNT -out $OUT_DIR

# ---------- АНОМАЛИЯ: WRONG_HEADING ----------
echo "📌 Генерация маршрутов с аномалией: wrong_heading..."
go run . \
  -fromX $FROMX -fromY $FROMY -toX $TOX -toY $TOY \
  -points $POINTS -type abnormal -anom wrong_heading -count $COUNT -out $OUT_DIR

# ---------- АНОМАЛИЯ: LOST_SIGNAL ----------
echo "📌 Генерация маршрутов с аномалией: lost_signal..."
go run . \
  -fromX $FROMX -fromY $FROMY -toX $TOX -toY $TOY \
  -points $POINTS -type abnormal -anom lost_signal -count $COUNT -out $OUT_DIR

# ---------- АНОМАЛИЯ: DEPTH_SPIKE ----------
echo "📌 Генерация маршрутов с аномалией: depth_spike..."
go run . \
  -fromX $FROMX -fromY $FROMY -toX $TOX -toY $TOY \
  -points $POINTS -type abnormal -anom depth_spike -count $COUNT -out $OUT_DIR

echo "Генерация завершена. Все CSV сохранены в $OUT_DIR"
