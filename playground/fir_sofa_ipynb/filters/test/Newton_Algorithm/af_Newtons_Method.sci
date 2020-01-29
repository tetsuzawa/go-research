//////////////////////////////////////////
//�@�@�@�K���t�B���^�i�j���[�g���@�j �@�@�@�@�@ 
//�@�@�@Adaptive Filter�iNewton's Method�j �@�@�@
//�@�@�@�@�@�@�@�@�@�@�@�@�@�@�@
//�@�@�@�@�@�@�@�@        �@�@�@M.Tsutsui 
////////////////////////////////////////// 

clear;

funcprot(0)
function [y_opt]=Newton(myu,Ite_C);//myu:�X�e�b�v�T�C�Y ,Ite_C:�X�V��

  for n=1:Ite_C;
       w_ini=(1-2*myu)*w_ini+2*myu*inv(sigma_x)*sigma_x_d;//�t�B���^�[�W���X�V_loop
  end

  y_opt=w_ini'*x;//�W���X�V��t�B���^�o��

endfunction


Data=50;//�f�[�^��
d=sin(1:0.2:Data)';//���͐M��
Data_renew=size(d,1);//�f�[�^�T�C�Y�X�V
w_ini=rand(Data_renew,Data_renew);//�K���t�B���^�����W��
v=2*rand(Data_renew,1)-2*rand(Data_renew,1);//�m�C�Y

x=d+v;//�K���t�B���^����

E_x=mean(x);//x����
E_d=mean(d);//d����
[L1,L2]=size(x);//�T�C�Y�`�F�b�N
sigma_x=1/L1*(x-E_x*ones(Data_renew,1))'*(x-E_x*ones(Data_renew,1));//x���ȑ���
sigma_x_d=1/L1*(x-E_x*ones(Data_renew,1))*(d-E_d*ones(Data_renew,1))';//x,d��

plot(d);//���]�M��
plot(x,'k');//�M��+�m�C�Y
plot(Newton(0.3,10),'r--');//�K���t�B���^�[�o��
xgrid();
title('Newtons Method');
legend(['Desired Signal';'Signal+Noise';'ADF Output';]);

