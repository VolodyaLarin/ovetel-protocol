import sys
import json
import bson


# Читаем JSON из стандартного ввода
data = sys.stdin.read()

try:
    json_data = json.loads(data)
except json.JSONDecodeError as e:
    print(f'Ошибка при парсинге JSON: {e}', file=sys.stderr)
    sys.exit(1)

bson_data = bson.dumps(json_data)

# Записываем результат в стандартный вывод
sys.stdout.buffer.write(bson_data)