# -*- coding: utf-8 -*-

# 打开输入文件和输出文件
with open('words_th.txt', 'r', encoding='utf-8') as input_file:
    with open('words_th_modified.txt', 'w', encoding='utf-8') as output_file:
        # 逐行读取输入文件
        for line in input_file:
            # 去除行尾的换行符
            line = line.strip()
            if len(line) > 0:
                # 在每行末尾添加 ",3,x"
                modified_line = f"{line}, 3, n\n"
                # 将修改后的行写入输出文件
                output_file.write(modified_line)

print("处理完成。修改后的内容已保存到 words_th_modified.txt 文件中。")

