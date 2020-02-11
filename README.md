# graduation-research

5年の卒業研究に使用したリポジトリ

卒業論文（[Raspbery Piを使用したドローンの駆動音低減法の検討](https://github.com/tetsuzawa/graduation-research/files/4184001/_.pdf)）

論文に載っているプログラムは[research-tools](https://github.com/tetsuzawa/research-tools)に置いてあります。

# directory structure

- documents  
  引き継ぎ資料や回路図
- paper  
  論文や発表予稿
- playground  
  実験的に作成したプログラムの置き場（研究と関係なし）
- slides  
  パワーポイントで作成した発表資料
- space-academia  
  スペースアカデミア関係の資料

- simulations
    - automatic_equalizer  
        - run_adf_x_dr_d_dr_voice_online  
            x: 参照信号、d: 目的信号 論文と同じ  
            ドローンと音声を伝達関数を畳み込んで合成したものに対する実験
        - run_adf_x_white_d_dr_voice_online  
            x: 白色雑音、d: ドローンの生の音  
            自動等化器の実験を行ったときに使用した
        - run_adf_x_white_d_dr_voice  
            上と同じだが、途中でプログラムが落ちるので、使わなかった
            
        - explore_mu  
            ドローンの駆動音に対する最適なステップサイズパラメータの探索を行った
        - explore_mu  
            白色雑音に対する最適なステップサイズパラメータの探索を行った
            
        - csvfiles  
            - auto  
                run_adf_x_white_d_dr_voiceを使用した自動等化器の実験、runのversion
            - auto_on  
                run_adf_x_white_d_dr_voice_onlinejを使用した自動等化器の実験、adaptのversion
            - auto_on_ref  
                run_adf_x_dr_d_dr_voice_onlineを使用した適応ノイズキャンセラの実験。
                確か参照信号に生のドローン、目的信号に生のドローンに白色信号を加えて実験したもの
            - auto_on_ref_convo  
                run_adf_x_dr_d_dr_voice_onlineを使用した適応ノイズキャンセラの実験、論文の第6章のもの。
                確か参照信号に生のドローン、目的信号に生のドローンに白色信号を加えて実験したもの
                
            - static  
                ドローンの駆動音に対する有色性の収束特性の実験結果
            - white  
                ドローンの駆動音に対する有色性の収束特性の実験結果
                
    - raw_drone_convergence  
        ドローンの駆動音に対する有色性の収束特性の実験に使用したはず
