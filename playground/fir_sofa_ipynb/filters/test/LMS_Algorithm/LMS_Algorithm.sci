/////////////////////////////////////////
//�@�@�@�K���t�B���^�iLMS�A���S���Y���j �@�@�@�@�@ 
//�@�@�@Adaptive Filter�iLMS Algorithm�j �@�@�@
//�@�@�@�@�@�@�@�@�@�@�@�@�@�@�@
//�@�@�@�@�@�@�@�@        �@�@M.Tsutsui 
//////////////////////////////////////// 

clear all;


funcprot(0)
function [y_opt]=lms_v(myu,update);//myu:�X�e�b�v�T�C�Y,update:�X�V��


  R=x*x'; //E[x�Ex�f]
  p=d*x; //E[d�Ex];

 
  buf=[];//�o�b�t�@�p��(�����ω��m�F)
  for n=1:update-1;
           w_renew=(eye(c_size,c_size)-myu*R)*w_renew+myu*p;//�t�B���^�[�W���X�V
          buf=[buf,w_renew];
  end

  y_opt=[]; 
  for i=1:1:update-1;
       y_opt=[y_opt,buf(:,i)'*x];//�W���X�V��t�B���^�o��
  end

endfunction

c_size=2;//�W���T�C�Y

d=3;//���]�T���v���l

x=rand(c_size,1);//���͐M��

w_renew=rand(c_size,1);//�K���t�B���^�����W��

y_opt_ini=w_renew'*x;//����o��



//__plot_______________
x_a=linspace(0,30,30);//0����`��
plot(x_a,[y_opt_ini,lms_v(0.1,30)],'b^--');
plot(x_a,[y_opt_ini,lms_v(0.2,30)],'k+--');
plot(x_a,[y_opt_ini,lms_v(0.3,30)],'ms--');
plot(x_a,[y_opt_ini,lms_v(0.4,30)],'rd--');
xgrid();
xlabel("�X�V��", "fontsize", 4);
legend(['��=0.1';'��=0.2';'��=0.3';'��=0.4';'��=0.5']);
title('LMS Algorithm','fontsize',5);
