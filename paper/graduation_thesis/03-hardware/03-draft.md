# ハードウェアの製作

この章では本研究で使用するドローン、マイク、Raspberry Piといったハードウェアの構成と製作について述べる。

## システムの外観


## ドローンについて

### 使用機器

ドローンはリンクスモーション株式会社の組み立て式ドローン
(Hquad500) を使用した。このドローンは機体にカーボンファイバーを採用しており、軽量、高強度、高剛性を兼ね備えている。また、拡張性が高く､ 容易に機能の追加が可能で様々な研究用途に適した製品である。

使用機器を以下に示す。

1. HQuad500 機体 [HQuad500](http://www.lynxmotion.com/p-1058-hquad500-hardware-only-kit.aspx)
   ![HQuad500](figures/hquad500_hardware.jpg)
   ![HQuad500](figures/hquad500_parts.jpg)
2. ESC (Electronic Speed Controller) [ESC](http://www.lynxmotion.com/p-915-12a-esc-simonk.aspx)
   ![ESC](figures/esc.jpg)


### 組み立ておよび動作確認

ドローンの組み立ては次のように行った。

1. 内容物の確認
   - 表でリスト
   - 写真
2. 機体の組み立て
3. ブラシレスモータ、ESC、フライトコントローラ、レシーバ、電源の配線
   - ブロック図
   - 写真
4. 動作確認


## バイノーラルマイクの製作

### マイクについて

収音には、音像定位実験のために製作したバイノーラルマイクを流用した。

使用した素子は秋月電子のエレクトレットコンデンサーマイクロホン（ECM）[XCM6035](http://akizukidenshi.com/catalog/g/gP-08181/)である。製作は@のようにLRそれぞれエレクトットコンデンサマイクをはんだ付けし、ロボットケーブルを経由してステレオミニプラグと接続した。

@TODO 価格を書くかどうか

- 使用機器
   1. エレクトットコンデンサマイク XCM6035 株式会社秋月電子通商 [url](http://akizukidenshi.com/catalog/g/gP-08181/) x2
   ![エレクトットコンデンサマイク](figures/microphone.jpg)
   ![エレクトットコンデンサマイク](figures/microphone_size.jpg)
      
   2. シールドスリムロボットケーブル KRT-SW 株式会社秋月電子通商 [url](http://akizukidenshi.com/catalog/g/gP-07457/)
   ![シールドスリムロボットケーブル](figures/sielded_robot_cable.jpg)
   ![シールドスリムロボットケーブル](figures/sielded_robot_cable_size.jpg)

   3. 3.5mmΦステレオミニプラグ MP-319 株式会社秋月電子通商
   ![3.5mmΦステレオミニプラグ](figures/mini_plug.jpg)
   ![3.5mmΦステレオミニプラグ](figures/mini_plug_size.jpg)

![バイノーラルマイクの製作](@TODO)


### アンプについて

前述のバイノーラルマイクのみでは十分な入力電圧が得られないため、収音にはマイクロフォンアンプを使用する必要がある。本研究では株式会社オーディオテクニカのマイクロフォンアンプAT-MA2をLRそれぞれ1台づつ使用した。

使用機器を以下に示す

- 使用機器
   1. マイクロホンアンプ AT-MA2 株式会社オーディオテクニカ [at-ma2](https://www.audio-technica.co.jp/product/AT-MA2)
   ![at-ma2](figures/at-ma2.jpg)
   2. 

なお、本研究では室内でサンプル収音を行ったため、据え置き型のマイクロフォンアンプを使用したが、実際にドローンに搭載する際は別途小型のマイクロフォンアンプが必要と推測される。


## Raspberry Piについて

Raspberry Piは英国のラズベリーパイ財団によって開発されている、ARMプロセッサを搭載したシングルボードコンピュータである。Raspberr Piは教育用として制作されたが、現在ではIoT製品開発などの業務や人工衛星のOBC (On Board Computer) にも使用されている。

Raspberry Piには

### OSの選定




