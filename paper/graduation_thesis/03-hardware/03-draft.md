# ハードウェアの製作

この章では本研究で使用するドローン、マイク、Raspberry Piといったハードウェアの構成と製作について述べる。

## システムの外観

まず、ドローンの本体のシステムについて述べる。


## ドローンについて

ドローンはリンクスモーション株式会社の組み立て式ドローン
(Hquad500) を使用した。このドローンは機体にカーボンファイバーを採用しており、軽量、高強度、高剛性を兼ね備えている。また、拡張性が高く､ 容易に機能の追加が可能で様々な研究用途に適した製品である。

### 使用機器

使用機器を以下に示す。

1. ドローンの機体 HQuad500 Hardware kit Lynxmotion株式会社 [HQuad500](http://www.lynxmotion.com/p-1058-hquad500-hardware-only-kit.aspx)
   ![HQuad500](figures/hquad500_hardware.jpg)
   ![HQuad500](figures/hquad500_parts.jpg)

2. ESC (Electronic Speed Controller) 12A ESC (SimonK) Lynxmotion株式会社 [ESC](http://www.lynxmotion.com/p-915-12a-esc-simonk.aspx)
   ![ESC](figures/esc.jpg)

3. ブラシレスモーター Brushless Motor 28x30 1000kv Lynxmotion株式会社 [ブラシレスモーター](http://www.lynxmotion.com/p-913-brushless-motor-28x30-1000kv.aspx)]
   ![ブラシレスモーター](figures/brushless_motor.jpg)

4. フライトコントローラー Quadrino Nano Lynxmotion株式会社  [フライトコントローラー](http://www.lynxmotion.com/p-1020-lynxmotion-quadrino-nano-flight-controller-with-gps.aspx)
   ![フライトコントローラー](figures/quadrino_nano.jpg)

5. リポバッテリー充電器 18W LiPo Battery Charger Lynxmotion株式会社 [リポバッテリー充電器](http://www.lynxmotion.com/p-985-18w-lipo-battery-charger.aspx)
   ![リポバッテリー充電器](figures/lipo_charger.jpg)

1. リポバッテリー 11.1V (3S), 3500mAh 30C LiPo Battery Pack
 Lynxmotion株式会社 [リポバッテリー充電器](http://www.lynxmotion.com/p-985-18w-lipo-battery-charger.aspx)
   ![リポバッテリー充電器](figures/lipo_charger.jpg)

1. ラジオレシーバー R9DS 10 channels 2.4GHz DSSS FHSS Receiver RadioLink株式会社 [ラジオレシーバー](http://www.radiolink.com.cn/doce/product-detail-120.html)
   ![ラジオレシーバー](figures/r9ds.jpg)

1. トランスミッタ AT9S 2.4GHz 10CH transmitter RadioLink株式会社 [トランスミッタ](http://www.radiolink.com.cn/doce/product-detail-119.html)
   ![トランスミッタ](figures/at9s.jpg)


### 組み立ておよび動作確認

ドローンの組み立ては@に従い、次のように行った。

1. 内容物の確認
   - 表でリスト
   - 写真
      ![drone_parts](drone_parts.png)

2. 機体の組み立て・源および信号線の配線
   ![drone_block](drone_block.pdf)
   ![drone_bloc](drone_block.png)

4. 動作確認
   [instructin](http://www.lynxmotion.com/images/document/PDF/LynxmotionUAV-QuadrinoNano-UserGuideV1.1.pdf)
   に従い、動作確認を行った。



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

- Linux系のOSで動作するためソフトウェア開発に強みをもち、GPIOピンを通してSPI、I2C、I2Sなどの通信を行えるため、センサなどを用いた開発を容易に行える。また、USB端子を搭載し、Wi-Fi、Bluetooth接続も可能で プロタイプ開発に適したデバイスとなっている。

### OSの選定

Raspberry Piで使用可能なOSには

- 電子工作などに適した公式OS Raspbian
- LinuxディストリビューションのUbuntuから派生した Ubuntu MATE
- Microsoft Windows 10

などが存在する。

本研究では主にGPIOを使用して開発を行うため、Raspbianを使用した。なお、バージョンは@TODO

### 初期設定について

1. OSのインストール
2. 地域、言語の設定
   `sudo raspi-config`  
   Localization Options

3. sshの設定
   `sudo raspi-config`  
   Interfacing Options  
   SSH  

4. プロキシに関する設定

3. アップデート

   ```shell
   sudo apt update 
   sudo apt upgrade -y 
   sudo apt dist-upgrade
   sudo rpi-update
   sudo reboot
   ```


- @TODO セキュリティに関する設定は省略

### AD変換用の拡張ボードについて

Raspberry PiはADC（ADコンバータ）を搭載していないため、マイクからの入力信号を扱うにはADコンバータを導入する必要がある。  

本研究で用いたのはマルツエレック株式会社の[Pumpkin Pi](http://select.marutsu.co.jp/list/detail.php?id=258)である。Pumpkin Piは計測用とオーディオ用のデュアルA-Dコンバータを搭載しており、Raspberry Piにオーディオ入力、アナログ入力機能を加えることが可能となる。

Pumpkin Piの仕様を以下に示す。

- 対応OS Raspbian
- 対応機種	Raspberry Pi Model B+/Raspberry Pi 2 Model B/Raspberry Pi 3 Model B
- LED出力	1点
- 赤外線リモコン機能	送受信
- オーディオコネクタ	φ3.5mmステレオミニジャック
- オーディオ入力	量子化ビット数=24,サンプリング周波数=48/96KHz
- 計測用AD変換	2チャンネル,16ビット
- 本体寸法	65(W)×56(D)mm
- 本体重量	約25g

![PumpkinPi](figures/pumpkin_pi.jpg)

[PumpkinPi](http://select.marutsu.co.jp/list/detail.php?id=258)
[PumpkinPi](https://www.marutsu.co.jp/pc/i/833515/)

### セットアップ

Pumpkin Piのセットアップは[トランジスタ技術 2017年1月号 オールDIPで1日製作!　音声認識ハイレゾPiレコーダ「Pumpkin Pi」](https://toragi.cqpub.co.jp/tabid/829/Default.aspx)
にしたがって行った。以下に簡易的な手順を示す。

1. Pumpkin Piを使用するためのRaspberry Pi固有の設定

   まず適当な作業ディレクトリで以下のコマンドを実行する。  

   ```shell
   wget http://einstlab.web.fc2.com/RaspberryPi/PumpkinPi.tar
   tar xvf PumpkinPi.tar
   cd PumpkinPi
   ./setup.sh  # @参考文献では ./PumpkinPi.sh と表記されている
   ```

2. カーネルとデバイス・ドライバのバージョンの確認  

   カーネルのバージョンとデバイス・ドライバのバージョンは同じである必要がある。カーネルのバージョンは`uname -r`で、デバイス・ドライバのバージョンは`modinfo snd_soc_pcm1808_adc.ko`でそれぞれ確認できる。

3. ADコンバータ用のデバイス・ドライバのインストール

   次の2つのデバイス・ドライバをインストールする。

   1. pcm1808-adc.ko  
      PCM1808固有の動作を決定するドライバ。 
   2. snd_soc_pcm1808_adc.ko  
      Raspberry Piのサウンドとして属性を決定するドライバ 

   まず、ホームディレクトリにPumpkinPi.tarをダウンロードして展開する。

   ```shell
   cd 
   wget http://einstlab.web.fc2.com/RaspberryPi/PumpkinPi.tar
   tar xvf PumpkinPi.tar
   cd PumpkinPi/Driver
   ```

   次にデバイス・ドライバをインストールする。

   ```shell
   sudo cp Backup/pcm1808-adc.bak/ /lib/modules/`uname -r`/kernel/sound/soc/codecs/pcm1808-adc.ko
   sudo cp Backup/snd_soc_pcm1808_adc.bak /lib/modules/`uname -r`/kernel/sound/soc/bcm/snd_soc_pcm1808_adc.ko
   sudo depmod -a  # 依存関係を調整
   ```

   OSのカーネル4.4以降ではデバイス・ツリー構造を導入してあるため、デバイス・ツリー情報ファイルをコピーする。

   ```shell
   sudo cp pcm1808-adc.dtbo /boot/oberlays/
   ```

   最後にデバイス・ドライバが電源起動時に自動的に読み込まれるように`/boot/config.txt`につぎの1行を追加する。

   ```shell
   dtoverlay=pcm1808-adc
   ```

   以上の作業を完了した後、再起動することで設定が適用される。








 








