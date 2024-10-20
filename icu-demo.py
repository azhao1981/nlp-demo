# -*- coding: utf-8 -*-
import icu

# 创建泰语分词器
bd = icu.BreakIterator.createWordInstance(icu.Locale('th'))

# 泰语文本
# สำหรับ
# หลาย
# คน
# แล้ว
text = "สำหรับหลายคนแล้ว สวัสดี พบกันพรุ่งนี้ มีของปีที่แล้ว ไปเมื่อเดือนที่แล้ว บ่ายห้าโมง เขตไห่เตี้ยน กรุงปักกิ่ง 13888886666"
text += "หนึ่งตัวสามอัน ทั้งหมดเป็นของคุณหวัง ต้องรออีกครึ่งวัน"

# 设置文本
bd.setText(text)

# 执行分词
start = bd.first()
for end in bd:
    print(text[start:end])
    start = end
