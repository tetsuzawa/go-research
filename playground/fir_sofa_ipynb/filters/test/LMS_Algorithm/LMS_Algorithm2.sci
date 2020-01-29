/////////////////////////////////////////
//�@�@�@�K���t�B���^�iLMS�A���S���Y���j �@�@�@�@�@ 
//�@�@�@Adaptive Filter�iLMS Algorithm�j �@�@�@
//�@�@�@�@�@�@�@�@�@�@�@�@�@�@�@
//�@�@�@�@�@�@�@�@        �@�@M.Tsutsui 
//////////////////////////////////////// 

clear;

c_size=3;//�W���T�C�Y
d=sin(0:0.5:100)';//���]�T���v�������]�M��

[size_r,size_r2]=size(d);//�T�C�Y�X�V

x=0.1*rand(size_r,1);//���͐M��

w=0.3*rand(size_r,1);//�K���t�B���^�W��

R=x*x'; //E[x�Ex�f]

funcprot(0)
function [y_opt]=lms_agm(myu,update);//myu:�X�e�b�v�T�C�Y,update:�X�V��


  p_buf=[];//p �T���v��d�ω�
  for i=1:1:size_r;
       p_buf=[p_buf,d(i,1)*x];
  end


  w_buf=[];//�W���o�b�t�@
  for p_loop=1:1:size_r; 
        for n=1:update-1;
             w=(eye(size_r,size_r)-myu*R)*w+myu*p_buf(:,p_loop);//�t�B���^�W���X�V
        end
        w_buf=[w_buf,w];
  end


  y_opt=w_buf'*x;//�K���t�B���^�o��

endfunction

//__plot_______________

subplot(3,1,1)
plot(d);
plot(lms_agm(0.1,5),'r--');
xgrid();
legend(['desired signal';'LMS']);
title('�X�V�� 3 ','fontsize',3);

subplot(3,1,2)
plot(d);
plot(lms_agm(0.1,10),'r--');
xgrid();
legend(['desired signal';'LMS']);
title('�X�V�� 5','fontsize',3);

subplot(3,1,3)
plot(d);
plot(lms_agm(0.1,15),'r--');
xgrid();
legend(['desired signal';'LMS']);
title('�X�V�� 10','fontsize',3);
