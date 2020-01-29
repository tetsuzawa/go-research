///////////////////////////////////////
//�@�@�@�K���t�B���^�i�ċA�ŏ����@�j �@�@�@�@�@ 
//�@�@�@Adaptive Filter
//�@�@�@Recursive least squares
//�@�@�@�@�@�@�@�@�@�@�@�@
//�@�@�@�@�@�@�@�@      �@�@�@M.Tsutsui 
///////////////////////////////////////

clear;

funcprot(0);
function[y_opt_buf]=RLS(arufa,lambda,update);//arufa:�� ,lambda:�Y�p�W�� ,update:�X�V��

 
     y_opt_buf=[];//ADF�o�̓o�b�t�@
 
     P_ini=1/arufa*eye(size_r2,size_r2);//P�����l


     for s_loop=1:1:size_r2;//__�T���v���ω�
            for i=1:1:update;//__�W���X�V���[�v__
        
                 gain=(1/lambda*(P_ini*x))/(1+1/lambda*x'*P_ini*x);//�Q�C���x�N�g��
   
                 e=d(1,s_loop)-w_ini'*x;//�덷 �T���v�����[�v
 
                 w_ini=w_ini+gain*e;//�W���X�V

               �@P_ini=1/lambda*(eye(size_r2,size_r2)-gain*x')*P_ini;//���ȑ��֍s��X�V   

            end//_______________�W���X�V���[�v__
        
             y_opt=w_ini'*x;//ADF�o�� 1Sample
             y_opt_buf=[y_opt_buf,y_opt];//ADF�o�̓x�N�g��
      
     end//_____________�T���v���ω�___

endfunction


d_size=80;//�f�[�^�T�C�Y

d=rand(0:1:d_size);//���]�M��

[size_r,size_r2]=size(d);//�T�C�Y�X�V

w_ini=zeros(size_r2,1);//�K���t�B���^�����W��

x=rand(size_r2,1);//�t�B���^����

plot(d);
plot(RLS(0.01,0.4,5),'r--');
xgrid();
title('�ċA�ŏ����@','fontsize',4);
legend(['Desired Signal';'RLS']);
