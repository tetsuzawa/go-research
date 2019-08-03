/////////////////////////////////////////
//　　　適応フィルタ（LMSアルゴリズム） 　　　　　 
//　　　Adaptive Filter（LMS Algorithm） 　　　
//　　　　　　　　　　　　　　　
//　　　　　　　　        　　M.Tsutsui 
//////////////////////////////////////// 

clear;

c_size=3;//係数サイズ
d=sin(0:0.5:100)';//所望サンプル→所望信号

[size_r,size_r2]=size(d);//サイズ更新

x=0.1*rand(size_r,1);//入力信号

w=0.3*rand(size_r,1);//適応フィルタ係数

R=x*x'; //E[x・x’]

funcprot(0)
function [y_opt]=lms_agm(myu,update);//myu:ステップサイズ,update:更新回数


  p_buf=[];//p サンプルd変化
  for i=1:1:size_r;
       p_buf=[p_buf,d(i,1)*x];
  end


  w_buf=[];//係数バッファ
  for p_loop=1:1:size_r; 
        for n=1:update-1;
             w=(eye(size_r,size_r)-myu*R)*w+myu*p_buf(:,p_loop);//フィルタ係数更新
        end
        w_buf=[w_buf,w];
  end


  y_opt=w_buf'*x;//適応フィルタ出力

endfunction

//__plot_______________

subplot(3,1,1)
plot(d);
plot(lms_agm(0.1,5),'r--');
xgrid();
legend(['desired signal';'LMS']);
title('更新回数 3 ','fontsize',3);

subplot(3,1,2)
plot(d);
plot(lms_agm(0.1,10),'r--');
xgrid();
legend(['desired signal';'LMS']);
title('更新回数 5','fontsize',3);

subplot(3,1,3)
plot(d);
plot(lms_agm(0.1,15),'r--');
xgrid();
legend(['desired signal';'LMS']);
title('更新回数 10','fontsize',3);
