/////////////////////////////////////
//�@�@�@�K���t�B���^�@�@�@�@�@ 
//�@�@�@LMS�A���S���Y���i�I�����C���j 
//�@�@�@Adaptive Filter
//�@�@�@�iLMS Algorithm(Online�j
//�@�@�@�@�@�@�@�@�@�@�@�@�@�@�@
//�@�@�@�@�@�@�@�@        �@M.Tsutsui 
//////////////////////////////////////

clear all;

d_size=80;//�f�[�^�T�C�Y
d=rand(0:1:d_size)';//���]�M��

[size_r,size_r2]=size(d);//�T�C�Y���X�V

x=rand(size_r,1);//���͐M��

w=ones(size_r,1);//�K���t�B���^�����W��


funcprot(0);
//__________________�֐�_______________________
//
// myu:�X�e�b�v�T�C�Y,update:�X�V��
//
// e_con:�X�V�I�������̐��l,y_opt:�K���t�B���^�o��
//______________________________________________

function[y_opt]=lms_opt_cef(myu,update,e_con);

  w_buf=[];//�W���o�b�t�@
  for s_loop=1:1:size_r;//�T���v���ω�
       for i=1:1:update;//__�W���X�V���[�v__
              y=w'*x;
              e=d(s_loop,1)-y;
              w=w+myu*e*x;  
              if (abs(e)<e_con) then,
                   w_buf=[w_buf,w];
                   break;
              end 
        end//_______________�W���X�V���[�v__
   end
  
   y_opt=w_buf'*x;
  
endfunction


//_plot_________________________
subplot(3,1,1)
plot(d);
plot(lms_opt_cef(0.02,50,2),'r--');
xgrid();
legend(['desired signal';'LMS']);
title('��:0.02,�X�V��:50,e:2','fontsize',3);

subplot(3,1,2)
plot(d);
plot(lms_opt_cef(0.05,50,1),'r--');
xgrid();
legend(['desired signal';'LMS']);
title('��:0.05,�X�V��:50,e:1','fontsize',3);

subplot(3,1,3)
plot(d);
plot(lms_opt_cef(0.02,50,0.1),'r--');
xgrid();
legend(['desired signal';'LMS']);
title('��:0.02,�X�V��:50,e:0.1','fontsize',3);

