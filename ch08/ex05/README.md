## 計測結果
コア数に合わせたほうが早い結果になっている。
CPUの計算の特性による理由？

Intel(R) Core(TM) i5-4258U CPU (2-core 4-thread)
```
processor = 1, time = 35.215388315s
processor = 2, time = 18.412155801s
processor = 3, time = 20.26812521s
processor = 4, time = 21.515716304s
processor = 5, time = 20.482134393s
processor = 6, time = 21.498821755s
processor = 7, time = 22.11240504s
processor = 8, time = 22.359376404s
```

デスクトップ用プロセッサでも試した結果、こちらはスレッド対応数だけ早くなっている。

Intel(R) Core(TM) i7-2600 CPU (4-core 8-thread)
```
processor = 1, time = 27.128667s
processor = 2, time = 28.5090077s
processor = 3, time = 22.9676403s
processor = 4, time = 22.9116332s
processor = 5, time = 11.8129041s
processor = 6, time = 10.2946862s
processor = 7, time = 9.9185478s
processor = 8, time = 9.7304223s
```
