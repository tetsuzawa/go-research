///////////////////////////////////////
//　　　適応フィルタ（再帰最小二乗法） 　　　　　 
//　　　Adaptive Filter
//　　　Recursive least squares
//　　　　　　　　　　　　
//　　　　　　　　      　　　M.Tsutsui 
///////////////////////////////////////

clear;

funcprot(0);
function[y_opt_buf]=RLS(arufa,lambda,update);//arufa:α ,lambda:忘却係数 ,update:更新回数

 
     y_opt_buf=[];//ADF出力バッファ
 
     P_ini=1/arufa*eye(size_r2,size_r2);//P初期値


     for s_loop=1:1:size_r2;//__サンプル変化
            for i=1:1:update;//__係数更新ループ__
        
                 gain=(1/lambda*(P_ini*x))/(1+1/lambda*x'*P_ini*x);//ゲインベクトル
   
                 e=d(1,s_loop)-w_ini'*x;//誤差 サンプルループ
 
                 w_ini=w_ini+gain*e;//係数更新

               　P_ini=1/lambda*(eye(size_r2,size_r2)-gain*x')*P_ini;//自己相関行列更新   

            end//_______________係数更新ループ__
        
             y_opt=w_ini'*x;//ADF出力 1Sample
             y_opt_buf=[y_opt_buf,y_opt];//ADF出力ベクトル
      
     end//_____________サンプル変化___

endfunction


d_size=80;//データサイズ

d=rand(0:1:d_size);//所望信号

[size_r,size_r2]=size(d);//サイズ更新

w_ini=zeros(size_r2,1);//適応フィルタ初期係数

x=rand(size_r2,1);//フィルタ入力

plot(d);
plot(RLS(0.01,0.4,5),'r--');
xgrid();
title('再帰最小二乗法','fontsize',4);
legend(['Desired Signal';'RLS']);
