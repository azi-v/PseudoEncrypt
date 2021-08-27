# PseudoEncrypt

<https://wiki.postgresql.org/wiki/Pseudo_encrypt>
<https://wiki.postgresql.org/wiki/XTEA_(crypt_64_bits)>
<https://wiki.postgresql.org/wiki/Pseudo_encrypt_constrained_to_an_arbitrary_range>

## pseudo_encrypt() 原理

生成指定范围的唯一整型id

```plpgsql
CREATE FUNCTION pseudo_encrypt_24(VALUE int) returns int AS $$
DECLARE
l1 int;
l2 int;
r1 int;
r2 int;
i int:=0;
BEGIN
 l1:= (VALUE >> 12) & (4096-1);
 r1:= VALUE & (4096-1);
 WHILE i < 3 LOOP
   l2 := r1;
   r2 := l1 # ((((1366 * r1 + 150889) % 714025) / 714025.0) * (4096-1))::int;
   l1 := l2;
   r1 := r2;
   i := i + 1;
 END LOOP;
 RETURN ((l1 << 12) + r1);
END;
$$ LANGUAGE plpgsql strict immutable;

CREATE FUNCTION bounded_pseudo_encrypt(VALUE int, max int) returns int AS $$
BEGIN
  loop
    value := pseudo_encrypt_24(value);
    exit when value <= max;
  end loop;
  return value;
END
$$ LANGUAGE plpgsql strict immutable;
```
