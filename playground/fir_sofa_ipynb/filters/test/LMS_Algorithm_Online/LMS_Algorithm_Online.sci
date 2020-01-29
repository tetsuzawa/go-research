/////////////////////////////////////
//　　　適応フィルタ　　　　　 
//　　　LMSアルゴリズム（オンライン） 
//　　　Adaptive Filter
//　　　（LMS Algorithm(Online）
//　　　　　　　　　　　　　　　
//　　　　　　　　        　M.Tsutsui 
//////////////////////////////////////

clear all;

d_size=80;//データサイズ
d=rand(0:1:d_size)';//所望信号

[size_r,size_r2]=size(d);//サイズ数更新

x=rand(size_r,1);//入力信号

w=ones(size_r,1);//適応フィルタ初期係数


funcprot(0);
//__________________関数_______________________
//
// myu:ステップサイズ,update:更新回数
//
// e_con:更新終了条件の数値,y_opt:適応フィルタ出力
//______________________________________________

function[y_opt]=lms_opt_cef(myu,update,e_con);

  w_buf=[];//係数バッファ
  for s_loop=1:1:size_r;//サンプル変化
       for i=1:1:update;//__係数更新ループ__
              y=w'*x;
              e=d(s_loop,1)-y;
              w=w+myu*e*x;  
              if (abs(e)<e_con) then,
                   w_buf=[w_buf,w];
                   break;
              end 
        end//_______________係数更新ループ__
   end
  
   y_opt=w_buf'*x;
  
endfunction


//_plot_________________________
subplot(3,1,1)
plot(d);
plot(lms_opt_cef(0.02,50,2),'r--');
xgrid();
legend(['desired signal';'LMS']);
title('μ:0.02,更新回数:50,e:2','fontsize',3);

subplot(3,1,2)
plot(d);
plot(lms_opt_cef(0.05,50,1),'r--');
xgrid();
legend(['desired signal';'LMS']);
title('μ:0.05,更新回数:50,e:1','fontsize',3);

subplot(3,1,3)
plot(d);
plot(lms_opt_cef(0.02,50,0.1),'r--');
xgrid();
legend(['desired signal';'LMS']);
title('μ:0.02,更新回数:50,e:0.1','fontsize',3);

