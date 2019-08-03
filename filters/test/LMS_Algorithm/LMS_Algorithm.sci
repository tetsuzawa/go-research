/////////////////////////////////////////
//　　　適応フィルタ（LMSアルゴリズム） 　　　　　 
//　　　Adaptive Filter（LMS Algorithm） 　　　
//　　　　　　　　　　　　　　　
//　　　　　　　　        　　M.Tsutsui 
//////////////////////////////////////// 

clear all;


funcprot(0)
function [y_opt]=lms_v(myu,update);//myu:ステップサイズ,update:更新回数


  R=x*x'; //E[x・x’]
  p=d*x; //E[d・x];

 
  buf=[];//バッファ用意(収束変化確認)
  for n=1:update-1;
           w_renew=(eye(c_size,c_size)-myu*R)*w_renew+myu*p;//フィルター係数更新
          buf=[buf,w_renew];
  end

  y_opt=[]; 
  for i=1:1:update-1;
       y_opt=[y_opt,buf(:,i)'*x];//係数更新後フィルタ出力
  end

endfunction

c_size=2;//係数サイズ

d=3;//所望サンプル値

x=rand(c_size,1);//入力信号

w_renew=rand(c_size,1);//適応フィルタ初期係数

y_opt_ini=w_renew'*x;//初回出力



//__plot_______________
x_a=linspace(0,30,30);//0から描画
plot(x_a,[y_opt_ini,lms_v(0.1,30)],'b^--');
plot(x_a,[y_opt_ini,lms_v(0.2,30)],'k+--');
plot(x_a,[y_opt_ini,lms_v(0.3,30)],'ms--');
plot(x_a,[y_opt_ini,lms_v(0.4,30)],'rd--');
xgrid();
xlabel("更新回数", "fontsize", 4);
legend(['μ=0.1';'μ=0.2';'μ=0.3';'μ=0.4';'μ=0.5']);
title('LMS Algorithm','fontsize',5);
