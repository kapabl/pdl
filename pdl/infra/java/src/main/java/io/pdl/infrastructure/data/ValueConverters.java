package io.pdl.infrastructure.data;

import java.nio.charset.StandardCharsets;
import java.sql.Time;
import java.sql.Timestamp;
import java.time.Instant;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;
import java.time.ZoneId;
import java.util.List;

final class ValueConverters {
    private static final ZoneId DEFAULT_ZONE = ZoneId.systemDefault();

    private ValueConverters() {
    }

    static Object convert(Object value, Class<?> targetType) {
        if (value == null) {
            return null;
        }
        if (targetType.isInstance(value)) {
            return value;
        }
        if (targetType == Integer.class || targetType == int.class) {
            return toInteger(value);
        }
        if (targetType == Long.class || targetType == long.class) {
            return toLong(value);
        }
        if (targetType == Double.class || targetType == double.class) {
            return toDouble(value);
        }
        if (targetType == Float.class || targetType == float.class) {
            return toFloat(value);
        }
        if (targetType == Short.class || targetType == short.class) {
            return toShort(value);
        }
        if (targetType == Byte.class || targetType == byte.class) {
            return toByte(value);
        }
        if (targetType == Boolean.class || targetType == boolean.class) {
            return toBoolean(value);
        }
        if (targetType == String.class) {
            return value.toString();
        }
        if (targetType == byte[].class) {
            return toBinary(value);
        }
        if (targetType == LocalDateTime.class) {
            return toLocalDateTime(value);
        }
        if (targetType == LocalDate.class) {
            return toLocalDate(value);
        }
        if (targetType == LocalTime.class) {
            return toLocalTime(value);
        }
        if (targetType == double[].class) {
            return toDoubleArray(value);
        }
        if (targetType == float[].class) {
            return toFloatArray(value);
        }
        return value;
    }

    private static Integer toInteger(Object value) {
        if (value instanceof Number number) {
            return number.intValue();
        }
        if (value instanceof Boolean bool) {
            return bool ? 1 : 0;
        }
        if (value instanceof String text) {
            String trimmed = text.trim();
            if (trimmed.isEmpty()) {
                return null;
            }
            return Integer.parseInt(trimmed);
        }
        return null;
    }

    private static Long toLong(Object value) {
        if (value instanceof Number number) {
            return number.longValue();
        }
        if (value instanceof Boolean bool) {
            return bool ? 1L : 0L;
        }
        if (value instanceof String text) {
            String trimmed = text.trim();
            if (trimmed.isEmpty()) {
                return null;
            }
            return Long.parseLong(trimmed);
        }
        return null;
    }

    private static Double toDouble(Object value) {
        if (value instanceof Number number) {
            return number.doubleValue();
        }
        if (value instanceof Boolean bool) {
            return bool ? 1d : 0d;
        }
        if (value instanceof String text) {
            String trimmed = text.trim();
            if (trimmed.isEmpty()) {
                return null;
            }
            return Double.parseDouble(trimmed);
        }
        return null;
    }

    private static Float toFloat(Object value) {
        if (value instanceof Number number) {
            return number.floatValue();
        }
        if (value instanceof Boolean bool) {
            return bool ? 1f : 0f;
        }
        if (value instanceof String text) {
            String trimmed = text.trim();
            if (trimmed.isEmpty()) {
                return null;
            }
            return Float.parseFloat(trimmed);
        }
        return null;
    }

    private static Short toShort(Object value) {
        if (value instanceof Number number) {
            return number.shortValue();
        }
        if (value instanceof Boolean bool) {
            return (short) (bool ? 1 : 0);
        }
        if (value instanceof String text) {
            String trimmed = text.trim();
            if (trimmed.isEmpty()) {
                return null;
            }
            return Short.parseShort(trimmed);
        }
        return null;
    }

    private static Byte toByte(Object value) {
        if (value instanceof Number number) {
            return number.byteValue();
        }
        if (value instanceof Boolean bool) {
            return (byte) (bool ? 1 : 0);
        }
        if (value instanceof String text) {
            String trimmed = text.trim();
            if (trimmed.isEmpty()) {
                return null;
            }
            return Byte.parseByte(trimmed);
        }
        return null;
    }

    private static Boolean toBoolean(Object value) {
        if (value instanceof Boolean bool) {
            return bool;
        }
        if (value instanceof Number number) {
            return number.intValue() != 0;
        }
        if (value instanceof String text) {
            String trimmed = text.trim();
            if (trimmed.isEmpty()) {
                return null;
            }
            return Boolean.parseBoolean(trimmed);
        }
        return null;
    }

    private static byte[] toBinary(Object value) {
        if (value instanceof byte[] data) {
            return data;
        }
        if (value instanceof String text) {
            return text.getBytes(StandardCharsets.UTF_8);
        }
        if (value instanceof List<?> list) {
            byte[] result = new byte[list.size()];
            for (int index = 0; index < list.size(); index++) {
                Object element = list.get(index);
                if (element instanceof Number number) {
                    result[index] = number.byteValue();
                } else if (element instanceof Boolean bool) {
                    result[index] = (byte) (bool ? 1 : 0);
                } else {
                    result[index] = 0;
                }
            }
            return result;
        }
        return null;
    }

    private static LocalDateTime toLocalDateTime(Object value) {
        if (value instanceof LocalDateTime dateTime) {
            return dateTime;
        }
        if (value instanceof Timestamp timestamp) {
            return timestamp.toLocalDateTime();
        }
        if (value instanceof java.util.Date date) {
            Instant instant = date.toInstant();
            return LocalDateTime.ofInstant(instant, DEFAULT_ZONE);
        }
        if (value instanceof LocalDate date) {
            return date.atStartOfDay();
        }
        if (value instanceof LocalTime time) {
            return LocalDate.now(DEFAULT_ZONE).atTime(time);
        }
        return null;
    }

    private static LocalDate toLocalDate(Object value) {
        if (value instanceof LocalDate date) {
            return date;
        }
        if (value instanceof java.sql.Date sqlDate) {
            return sqlDate.toLocalDate();
        }
        if (value instanceof Timestamp timestamp) {
            return timestamp.toInstant().atZone(DEFAULT_ZONE).toLocalDate();
        }
        if (value instanceof java.util.Date date) {
            return date.toInstant().atZone(DEFAULT_ZONE).toLocalDate();
        }
        if (value instanceof LocalDateTime dateTime) {
            return dateTime.toLocalDate();
        }
        return null;
    }

    private static LocalTime toLocalTime(Object value) {
        if (value instanceof LocalTime time) {
            return time;
        }
        if (value instanceof Time sqlTime) {
            return sqlTime.toLocalTime();
        }
        if (value instanceof Timestamp timestamp) {
            return timestamp.toInstant().atZone(DEFAULT_ZONE).toLocalTime();
        }
        if (value instanceof LocalDateTime dateTime) {
            return dateTime.toLocalTime();
        }
        return null;
    }

    private static double[] toDoubleArray(Object value) {
        if (value instanceof double[] array) {
            return array;
        }
        if (value instanceof Double[] boxed) {
            double[] result = new double[boxed.length];
            for (int index = 0; index < boxed.length; index++) {
                Double entry = boxed[index];
                result[index] = entry != null ? entry : 0d;
            }
            return result;
        }
        if (value instanceof List<?> list) {
            double[] result = new double[list.size()];
            for (int index = 0; index < list.size(); index++) {
                Object entry = list.get(index);
                if (entry instanceof Number number) {
                    result[index] = number.doubleValue();
                } else if (entry instanceof Boolean bool) {
                    result[index] = bool ? 1d : 0d;
                } else {
                    result[index] = 0d;
                }
            }
            return result;
        }
        if (value.getClass().isArray()) {
            Object[] array = (Object[]) value;
            double[] result = new double[array.length];
            for (int index = 0; index < array.length; index++) {
                Object entry = array[index];
                if (entry instanceof Number number) {
                    result[index] = number.doubleValue();
                } else if (entry instanceof Boolean bool) {
                    result[index] = bool ? 1d : 0d;
                } else {
                    result[index] = 0d;
                }
            }
            return result;
        }
        return null;
    }

    private static float[] toFloatArray(Object value) {
        if (value instanceof float[] array) {
            return array;
        }
        if (value instanceof Float[] boxed) {
            float[] result = new float[boxed.length];
            for (int index = 0; index < boxed.length; index++) {
                Float entry = boxed[index];
                result[index] = entry != null ? entry : 0f;
            }
            return result;
        }
        if (value instanceof List<?> list) {
            float[] result = new float[list.size()];
            for (int index = 0; index < list.size(); index++) {
                Object entry = list.get(index);
                if (entry instanceof Number number) {
                    result[index] = number.floatValue();
                } else if (entry instanceof Boolean bool) {
                    result[index] = bool ? 1f : 0f;
                } else {
                    result[index] = 0f;
                }
            }
            return result;
        }
        if (value.getClass().isArray()) {
            Object[] array = (Object[]) value;
            float[] result = new float[array.length];
            for (int index = 0; index < array.length; index++) {
                Object entry = array[index];
                if (entry instanceof Number number) {
                    result[index] = number.floatValue();
                } else if (entry instanceof Boolean bool) {
                    result[index] = bool ? 1f : 0f;
                } else {
                    result[index] = 0f;
                }
            }
            return result;
        }
        return null;
    }
}
