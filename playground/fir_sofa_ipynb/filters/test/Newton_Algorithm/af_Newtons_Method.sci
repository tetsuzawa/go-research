//////////////////////////////////////////
//　　　適応フィルタ（ニュートン法） 　　　　　 
//　　　Adaptive Filter（Newton's Method） 　　　
//　　　　　　　　　　　　　　　
//　　　　　　　　        　　　M.Tsutsui 
////////////////////////////////////////// 

clear;

funcprot(0)
function [y_opt]=Newton(myu,Ite_C);//myu:ステップサイズ ,Ite_C:更新回数

  for n=1:Ite_C;
       w_ini=(1-2*myu)*w_ini+2*myu*inv(sigma_x)*sigma_x_d;//フィルター係数更新_loop
  end

  y_opt=w_ini'*x;//係数更新後フィルタ出力

endfunction


Data=50;//データ数
d=sin(1:0.2:Data)';//入力信号
Data_renew=size(d,1);//データサイズ更新
w_ini=rand(Data_renew,Data_renew);//適応フィルタ初期係数
v=2*rand(Data_renew,1)-2*rand(Data_renew,1);//ノイズ

x=d+v;//適応フィルタ入力

E_x=mean(x);//x平均
E_d=mean(d);//d平均
[L1,L2]=size(x);//サイズチェック
sigma_x=1/L1*(x-E_x*ones(Data_renew,1))'*(x-E_x*ones(Data_renew,1));//x自己相関
sigma_x_d=1/L1*(x-E_x*ones(Data_renew,1))*(d-E_d*ones(Data_renew,1))';//x,d相

plot(d);//所望信号
plot(x,'k');//信号+ノイズ
plot(Newton(0.3,10),'r--');//適応フィルター出力
xgrid();
title('Newtons Method');
legend(['Desired Signal';'Signal+Noise';'ADF Output';]);

