# 10000000倍するとcomplex64の精度ではフラクタルを表現しきれない状態となりました。
# bigfloat と complex128 はほぼ同等の描写を実現していますが、bigfloatの計算には
# とても大きな時間がかかります。（下記参照）
# ratの計算には膨大な計算時間がかかるため、イテレーションを多く回すことができず、結果的に精度が
# 高くありません。

# 各実行時間
# complex128 = 0.419681 Seconds
# complex64 = 0.507608 Seconds
# big.Float = 25.983636 Seconds
# big.Rat(ite = 3) = 223.114646 Seconds
go run complex128/main.go > complex128.png
go run complex64/main.go > complex64.png

go run bigfloat/main.go > bigfloat.png
go run rat/main.go > rat.png
