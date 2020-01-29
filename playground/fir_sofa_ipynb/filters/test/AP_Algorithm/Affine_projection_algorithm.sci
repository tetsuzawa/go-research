///////////////////////////////////////
//�@�@�@�K���t�B���^�i�A�t�B���ˉe�@�j �@�@�@�@�@ 
//�@�@�@Adaptive Filter �@�@�@
//�@�@�@Affine projection algorithm�@�@�@�@�@�@�@�@�@�@�@�@
//�@�@�@�@�@�@�@�@         M.Tsutsui 
////////////////////////////////////// 

clear all;

funcprot(0)
 function [y_opt]=APA(myu,arufa,RC);
 //____________APA�@�֐� ___________//
 //
 // myu:�X�e�b�v�T�C�Y, arufa:���������̒萔(������)���t�s��̈��萫�m��
//
 // RC:�X�V��
//
 //____________APA�@�֐� ___________//


 for i=1:1:RC-1;
  
      e=d-w_ini'*X;//�덷�x�N�g��
 
      w_ini=w_ini+myu*X*inv(arufa*eye(size_r2,size_r2)+X'*X)*e';//���K��
 
  end

 y_opt=w_ini'*X;//�t�B���^�o��

endfunction


d_size=80;//�f�[�^�T�C�Y

d=rand(0:0.5:d_size);//���]�M���x�N�g��

[size_r,size_r2]=size(d);//�T�C�Y�X�V

w_ini=rand(size_r2,1);//�K���t�B���^�����W��

X=w_ini*d;//���̓x�N�g��


plot(d);
plot(APA(0.5,3,9),'r--');
xgrid();
legend(['desired signal';'APA']);
title('Affine projection algorithm','fontsize',4);
