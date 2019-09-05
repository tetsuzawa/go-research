///////////////////////////////////////
//　　　適応フィルタ（アフィン射影法） 　　　　　 
//　　　Adaptive Filter 　　　
//　　　Affine projection algorithm　　　　　　　　　　　　
//　　　　　　　　         M.Tsutsui 
////////////////////////////////////// 

clear all;

funcprot(0)
 function [y_opt]=APA(myu,arufa,RC);
 //____________APA法関数 ___________//
 //
 // myu:ステップサイズ, arufa:小さい正の定数(正則化)※逆行列の安定性確保
//
 // RC:更新回数
//
 //____________APA法関数 ___________//


 for i=1:1:RC-1;
  
      e=d-w_ini'*X;//誤差ベクトル
 
      w_ini=w_ini+myu*X*inv(arufa*eye(size_r2,size_r2)+X'*X)*e';//正規化
 
  end

 y_opt=w_ini'*X;//フィルタ出力

endfunction


d_size=80;//データサイズ

d=rand(0:0.5:d_size);//所望信号ベクトル

[size_r,size_r2]=size(d);//サイズ更新

w_ini=rand(size_r2,1);//適応フィルタ初期係数

X=w_ini*d;//入力ベクトル


plot(d);
plot(APA(0.5,3,9),'r--');
xgrid();
legend(['desired signal';'APA']);
title('Affine projection algorithm','fontsize',4);
